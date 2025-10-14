use v5.40;
use experimental qw<class>;

class CustomSet {

    field $elements : reader : param;

    method is_empty () {
        return scalar @{$elements} == 0;
    }

    method contains ($element) {
        foreach my $elem ( @{$elements} ) {
            return 1 if $element == $elem;
        }

        return 0;
    }

    method is_subset_of ($other) {
        for my $elem ( @{ $self->elements } ) {
            return 0 unless $other->contains($elem);
        }
        return 1;
    }

    method is_disjoint_of ($other) {
        for my $elem ( @{ $self->elements } ) {
            return 0 if $other->contains($elem);
        }
        for my $elem ( @{ $other->elements } ) {
            return 0 if $self->contains($elem);
        }
        return 1;
    }

    method is_equal_to ($other) {
        return $self->is_subset_of($other) && $other->is_subset_of($self);
    }

    method add ($element) {
        push @{ $self->elements }, $element;
        return $self;
    }

    method intersection ($other) {
        my @i = ();
        for my $elem ( @{ $self->elements } ) {
            push @i, $elem if $other->contains($elem);
        }

        return CustomSet->new( elements => \@i );
    }

    method difference ($other) {
        my @i = ();
        for my $elem ( @{ $self->elements } ) {
            push @i, $elem unless $other->contains($elem);
        }

        return CustomSet->new( elements => \@i );
    }

    method union ($other) {
        my @u = ( @{ $self->elements }, @{ $other->elements } );

        return CustomSet->new( elements => \@u );
    }
}

1;
