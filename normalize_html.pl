#!/usr/bin/env perl

use 5.036;

use FindBin;
use File::Find;
use File::Slurp;
use XML::LibXML;

find(
    sub {
        return unless $_ =~ /.*.html$/xmi;
        say "Processing: $_";
        my $loc  = "$FindBin::Bin/posts/$_";
        my $html = read_file($loc);
        my $doc  = XML::LibXML->new->parse_html_string($html);

        my ($html_tag) = $doc->getElementsByTagName('html');
        $html_tag->setAttribute( lang => 'en' );

        my @meta_tags = $doc->getElementsByTagName('meta');

        for my $meta (@meta_tags) {
            if ( my $v = $meta->getAttribute('value') ) {
                $meta->removeAttribute('value');
                $meta->setAttribute( content => $v );
            }
        }

        my @links = $doc->getElementsByTagName('link');

        for my $link (@links) {
            if ( $link->getAttribute('href') eq '/index.css' ) {
                say "Already done $_";
                return;
            }
        }

        my $st = XML::LibXML::Element->new('link');
        $st->setAttribute( rel  => 'stylesheet' );
        $st->setAttribute( href => '/index.css' );
        my ($head) = $doc->getElementsByTagName('head');
        $head->appendChild($st);
        write_file( $loc, $doc->toString() );
    },
    "$FindBin::Bin/posts"
);
