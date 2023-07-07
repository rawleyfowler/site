#!/usr/bin/env perl
use 5.016;

use Mojolicious::Lite -signatures;
use Mojo::SQLite;

my $sql = Mojo::SQLite->new('sqlite:site.db');
helper db => sub { state $db = $sql->db };

$sql->migrations->name('posts_table')->from_string(<<EOF)->migrate;
-- 1 up
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    content TEXT,
    slug TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- 1 down
DROP TABLE posts;
EOF

$sql->migrations->name('posts_index')->from_string(<<EOF)->migrate;
-- 2 up
CREATE INDEX IF NOT EXISTS slug_idx ON posts (slug);
-- 2 down
DROP INDEX slug_idx;
EOF

get '/' => sub {
    my $c = shift;
    $c->reply->static('index.html');
};

get '/blog' => sub {
    my $c = shift;
    $c->stash( posts => [ $c->db->select('posts')->hashes->each ] );
    $c->render;
};

get '/blog/:post' => sub {
    my $c    = shift;
    my $post = $c->db->select(
        'posts',
        [ 'title', 'content', 'slug' ],
        { slug => $c->param('post') }
    )->hash;

    return $c->render( template => 'not_found' ) unless $post;

    $c->stash( post => $post );
    $c->render( template => 'post' );
};

app->start;

__DATA__
@@ not_found.html.ep
  <!DOCTYPE html>
  <html lang="en" data-bs-theme="dark">
  %= include '_header', title => '404 Not Found'
  <body class="container fw-light">
  %= include '_nav'
  <h1 class="mt-3 display-2">404 - Not Found</h1>
  </body>
  </html>
  
@@ post.html.ep
  % use Text::Markdown qw(markdown);
  <!DOCTYPE html>
  <html lang="en" data-bs-theme="dark">
  %= include '_header', title => $post->{title}
  <body class="container fw-light">
  %= include '_nav'
  <h1 class="mt-3 display-2"><%= $post->{title} %></h1>
  <div>
  <%== markdown($post->{content}) %>
  </div>
  </body>
  </html>
  
@@ blog.html.ep
  <!DOCTYPE html>
  <html lang="en" data-bs-theme="dark">
  %= include '_header', title => 'Blog'
  <body class="container fw-light">
  %= include '_nav'
  <h1 class="mt-3 display-2">Blog</h1>
  <p>
  This is my blog; it represents my opinions, not that of my employer(s) past, present or future.
  My blog is made up of Functional Programming, Perl & Raku, Rants, Politics, and whatever else I dream up.
  If you are prone to becoming upset about things, this might not be the best place for you.
  </p>
  <p>
  All content is licensed under the <a href="https://creativecommons.org/licenses/by-sa/3.0/">Creative-Commons Attribution-ShareALike 3.0</a> license unless specified otherwise.
  </p>
  <div class="d-flex flex-row justify-content-between mt-3">
  <h2>Title</h2>
  <h2>Publish Date</h2>
  </div>
  % for (@$posts) {
    <div class="d-flex flex-row justify-content-between mt-3">
        <a href="/blog/<%= $_->{slug} %>"><%= $_->{title} %></a>
        <div><%= $_->{created_at} %></div>
    </div>
  % }
  </body>
  </html>

@@ _header.html.ep
  <head>
  <title><%= $title %></title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>
  <style>
  body * {
    font-size: calc(0.33vw + 12px);
  }
  </style>
  </head>

@@ _nav.html.ep
  <nav class="d-flex flex-row text-decoration-none mt-3">
  <div class="me-3">
  <a href="/">Home</a>
  </div>
  <div>
  <a href="/blog">Blog</a>
  </div>
  </nav>
