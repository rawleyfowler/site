open Lwt
open Views

module Blog_post = Database.Blog_post

module type SimpleRender = (sig val render : unit -> string end)

let header_template =
  {eos|<!DOCTYPE html>
    <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="Rawley.xyz, Rawley Fowler's blog and personal website">
    <title>rawley.xyz</title>
    <link href="/static/index.css" rel="stylesheet">
    <link rel="alternate" type="application/rss+xml" href="https://rawley.xyz/blog/rss.xml" title="rawley.xyz">
    </head>
    <body>
    <body>
    <header>
    <h1>
    <a href="/">Rawley.xyz</a>
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
    <main>|eos}

let footer_template =
  {eos|</main>
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
    </html>|eos}

let replace_sequence r t s =
  Str.(global_replace (regexp r) t s)

let html_unescape s =
  s
  |> replace_sequence "&lsquo;" "'"
  |> replace_sequence "&rsquo;" "'"
  |> replace_sequence "&gt;" ">"
  |> replace_sequence "&lt;" "<"

let not_found_template =
  Printf.sprintf "%s %s %s"
    header_template
    "<h4>404 Not found<h4>"
    footer_template

let error_template =
  Printf.sprintf "%s %s %s"
    header_template
    "<h4>500 Server Error</h4>"
    footer_template

let render_page content =
  Printf.sprintf
    "%s %s %s"
    header_template
    content
    footer_template

let generate_link (p : Blog_post.t) =
  Printf.sprintf
    {eos|<div class="link-wrapper">
      <a href="/blog/%s">%s</a>
      <i>%s</i>
      </div>|eos}
    p.slug p.title p.date

let generate_rss_item (p : Blog_post.t) =
  Printf.sprintf
    {eos|<item>
      <title>%s</title>
      <link>https://rawley.xyz/blog/%s</link>
      <description></description>
      </item>|eos}
    p.title p.slug

let handle_error e =
  print_endline (Caqti_error.show e);
  Dream.html ?code:(Some 500) error_template

let handle_not_found () =
  Dream.html ?code:(Some 404) not_found_template

let render_blog_index () =
  let%lwt posts = Database.get_all_blog_posts () in
  match posts with
  | Error e -> handle_error e
  | Ok posts -> Blog_index.render ~posts |> Layout.render |> Dream.html

let render_rss_feed (_ : Dream.request) =
  let buff = Buffer.create 512 in
  let add_str = Buffer.add_string buff in
  let () =
    add_str
      {eos|<?xml version="1.0" encoding="UTF-8" ?>
        <rss version="2.0">
        <channel>
        <title>rawley.xyz blog</title>
        <description>Functional programming, math, and philosophy</description>
        <image>
        <url>https://rawley.xyz/static/rawley.xyz.png</url>
        <link>https://rawley.xyz/</link>
        </image>|eos}
  in
  let posts_t = Database.get_all_blog_posts () in
  posts_t >>= function
  | Error e -> handle_error e
  | Ok posts ->
      let () = List.iter (fun t -> generate_rss_item t
                                  |> html_unescape
                                  |> add_str) posts
      in
      let () =
        add_str
          {eos|</channel>
          </rss>|eos}
      in
      Lwt.return @@
        Dream.response
          ~headers:["Content-Type", "text/xml";]
          (Buffer.contents buff)

let render_simple (module R : SimpleRender) =
  R.render () |> Layout.render |> Dream.html

let render_blog_post ~slug =
  let%lwt post = Database.get_blog_post_by_slug slug in
  match post with
  | Ok p_opt ->
    begin match p_opt with
    | None -> handle_not_found ()
    | Some p ->
      Dream.html @@ Post_layout.render ~title:p.title ~body:p.content ~tags:[p.title]
  end
  | Error e -> handle_error e


let render_index () =
  render_simple (module Index)
