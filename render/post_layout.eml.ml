let render ~title ~tags ~body =
  let combine_tags lst =
    List.fold_left (fun t a -> t ^ " " ^ a) "" lst
  in
  <html>
  <head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="<%s combine_tags tags %>">
  <title><%s title %></title>
  <link href="/static/index.css" rel="stylesheet">
  </head>
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
  <h2><%s title %></h2>
  <%s! body %>
  </main>
  <footer>
  <hr>
  <p>
  Copyright &copy; Rawley Fowler 2022
  </p>
  <p>
  <b>Disclaimer</b>: All opinions on this site, are that of my own. They do not reflect the opinions of 
  any of my employers; past, present, or future.
  </p>
  </footer>
  </body>
  </html>