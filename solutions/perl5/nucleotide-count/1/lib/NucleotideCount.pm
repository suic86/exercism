package NucleotideCount;

use strict;
use warnings;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<count_nucleotides>;

sub count_nucleotides ($strand) {
    die "Invalid nucleotide in strand" unless $strand =~ /^[ACGT]*$/;
    my %nucleotides = ( 'A' => 0, 'C' => 0, 'G' => 0, 'T' => 0 );
    for ( $strand =~ /[ACTG]/g ) {
        $nucleotides{$_} += 1;
    }
    return \%nucleotides;
}

1;
