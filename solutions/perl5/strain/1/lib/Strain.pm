package Strain;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<keep discard>;

sub keep ( $input, $function ) {
    return [ grep { $function->($_) } @$input ];
}

sub discard ( $input, $function ) {
    return [ grep { !$function->($_) } @$input ];
}

1;
