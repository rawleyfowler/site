#!/usr/bin/env perl

use strict;
use warnings;

use Carp qw(croak);
use Mojo::SQLite;
use Mojo::File;
use feature qw(say);

my $db = Mojo::SQLite->new('sqlite:site.db')->db;

my $title   = shift || croak 'Expected $1 to be title';
my $slug    = shift || croak 'Expected $2 to be slug';
my $content = shift || croak 'Expected $3 to be md file';

$content = Mojo::File->new($content)->slurp;

my $post = {
    title   => $title,
    slug    => $slug,
    content => $content
};

if ( my $old_post = $db->select( 'posts', ['id'], { slug => $slug } )->hash ) {
    say 'Updating post.';
    $db->update( 'posts', $post, { id => $old_post->{id} } );
}
else {
    say 'Inserting post.';
    $db->insert( 'posts', $post );
}

say 'Finished.';
