package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nomkhonwaan/myblog-server/protos/blog-service/posts"
	"google.golang.org/grpc"
)

var (
	blogservicePostsEndpoint   = flag.String("blogservice_posts_endpoint", os.Getenv("GRPC_SERVER_URI"), "endpoint of blogservice_posts.PostsService")
	blogservicePostsSwaggerDir = flag.String("blogservice_posts_swagger_dir", "protos/blog-service/posts", "path to the directory which contains swagger definitions")
)

func serveBlogServicePostsSwagger(rw http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		glog.Errorf("Not Found: %s", r.URL.Path)
		http.NotFound(rw, r)
		return
	}

	glog.Info("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/posts/")
	p = path.Join(*blogservicePostsSwaggerDir, p)
	http.ServeFile(rw, r, p)
}

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := protos_blogservice_posts.RegisterPostsServiceHandlerFromEndpoint(ctx, mux, *blogservicePostsEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

func Run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/posts/", serveBlogServicePostsSwagger)

	gw, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	http.ListenAndServe(address, allowCORS(mux))
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(":" + os.Getenv("PORT")); err != nil {
		glog.Fatal(err)
	}
}
