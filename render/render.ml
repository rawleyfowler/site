open Database
open Lwt

module Render = struct
  let not_found_template = "404 Not found"
  let error_template = "500 Server Error"

  let header_template =
    {eos|
     <head>
     <meta charset="utf-8">
     <meta name="viewport" content="width=device-width, initial-scale=1">
     <meta name="description" content="Rawley.xyz, Rawley Fowler's blog and personal website">
     <title>rawley.xyz</title>
     <link href="/static/index.css" rel="stylesheet">
     </head>
     <body>
     <body>
     <header>
     <h1>
     <a href="/">/home/rawley.xyz</a>
     </h1>
     <nav>
     <li>
     <a href="/blog">blog</a>
     </li>
     <li>
     <a href="/resume">resume</a>
     </li>
     <li>
     <a href="/philosophy">philosophy</a>
     </li>
     <li>
     <a href="/web-ring">web ring</a>
     </li>
     </nav>
     </header>
     <main>
     |eos}

  let footer_template =
    {eos|
     </main>
     <footer>
     <hr>
     <p>
     Copyright &copy; Rawley Fowler 2022
     </p>
     <p>
     <b>Disclaimer</b>: All opinions on this site, are that of my own. They do not
     reflect the opinions of any of my employers; past, present, or future.
     </p>
     </footer>
     </body>
     </html>
     |eos}

  let render_page content =
    Printf.sprintf
      {eos| 
       %s
       %s
       %s
       |eos}
      header_template
      content
      footer_template
    
  let generate_link (p : BlogPost.t) =
    Printf.sprintf
      {eos|
       <div class="link-wrapper">
       <a href="/blog/%s">%s</a>
       <i>%s</i>
       </div>
       |eos}
      p.slug p.title p.date
  
  let handle_error e =
    print_endline (Caqti_error.show e); error_template |> Dream.html

  let render_blog_post request =
    let slug = Dream.param request "post" in
    let post_t = Database.get_blog_post_by_slug slug in
    post_t >>= fun post ->
    match post with
    | Error e -> handle_error e
    | Ok p_opt ->
       match p_opt with
       | None -> not_found_template |> Dream.html
       | Some p ->
          Printf.sprintf "<h2>%s</h2>\r\n%s" p.title p.content
          |> render_page
          |> Dream.html

  let render_blog_index (_ : Dream.request) =
    let buff = Buffer.create 512 in
    let () = Buffer.add_string buff "<h3>Blog</h3>" in
    let posts_t = Database.get_all_blog_posts () in
    posts_t >>= function
    | Error e -> handle_error e
    | Ok posts ->
       List.iter (fun p -> Buffer.add_string buff (generate_link p)) posts;
       let c = if List.length posts <> 0 then
                 render_page @@ Buffer.contents buff
               else render_page "<br>No blog posts..." in
       Dream.html c
end
