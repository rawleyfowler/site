open Lwt.Infix
open Lwt.Syntax

let () = Lwt_main.run @@ Database.create_blog_post_table () >>= function
  | Ok () -> Lwt_io.print "Database initialized successfully.\n"
  | Error e -> failwith (Caqti_error.show e)

let () =
  Dream.run
  @@ Dream.logger
  @@ Dream.router [
         Dream.get "/static/**" @@ Dream.static "static";
         Dream.get "/favicon.ico" @@ Dream.from_filesystem "static" "favico.png";
         Dream.get "/" @@ (fun _ -> Render.render_index ());
         Dream.get "/resume" @@ Dream.from_filesystem "html" "resume.html";
         Dream.get "/philosophy" @@ Dream.from_filesystem "html" "philosophy.html";
         Dream.get "/web-ring" @@ Dream.from_filesystem "html" "web-ring.html";
         Dream.get "/blog" (fun _ -> Render.render_blog_index ());
         Dream.scope "/blog" [] [
             Dream.get "/rss.xml" Render.render_rss_feed;
             Dream.get "/:post" (fun request -> Render.render_blog_post ~slug:(Dream.param request "post"))
           ];
         Dream.get "/**" (fun _ -> Render.handle_not_found ())
       ]
