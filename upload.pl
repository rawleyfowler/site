#!/usr/bin/env perl

use strict;
use warnings;

use Mojo::UserAgent;
use Mojo::File;
use Carp qw(croak carp);

my $file = Mojo::File->new(shift)
  || croak q{Expected $1 to be a file path, or couldn't open file};

my $ua = Mojo::UserAgent->new;
my $hn = shift || croak q{Expected $2 to be header name};
my $pw = shift || croak q{Expected $3 to be password};
my $ti = shift || croak q{Expected $4 to be title};
my $sl = shift || croak q{Expected $5 to be slug};
my $uri =
  $ENV{BLOG_DEVELOPMENT} ? 'http://localhost:3000' : 'https://rawley.xyz';

my $j = {
    content => $file->slurp,
    title   => $ti,
    slug    => $sl
};

my $r =
  $ua->post( ( $uri . '/blog' ) => { $hn => $pw } => json => $j )->result;

if ( $r->is_success ) {
    print 'Upload successful';
}
else {
    carp q{Failed to upload } . $r->code;
    exit 1;
}

