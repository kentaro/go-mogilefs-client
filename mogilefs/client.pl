#!/usr/bin/env perl
use strict;
use warnings;

use MogileFS::Client;

my $mogfs = MogileFS::Client->new(
    domain => 'go-mogilefs-client',
    hosts  => ['127.0.0.1:7001'],
);

my $key = 'test';
my $fh = $mogfs->new_file($key, 'test') or die $mogfs->errstr;
$fh->print('hello');
$fh->close or die $mogfs->errstr;

my $data = $mogfs->get_file_data($key) or die $mogfs->errstr;
print $$data . "\n";

my @paths = $mogfs->get_paths($key);
print "$_\n" for @paths;

$mogfs->delete($key);
