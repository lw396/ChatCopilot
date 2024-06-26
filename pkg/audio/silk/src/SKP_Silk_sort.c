 

/* Insertion sort (fast for already almost sorted arrays):   */
/* Best case:  O(n)   for an already sorted array            */
/* Worst case: O(n^2) for an inversely sorted array          */

#include "SKP_Silk_SigProc_FIX.h"

void SKP_Silk_insertion_sort_increasing(
    SKP_int32           *a,             /* I/O:  Unsorted / Sorted vector               */
    SKP_int             *index,         /* O:    Index vector for the sorted elements   */
    const SKP_int       L,              /* I:    Vector length                          */
    const SKP_int       K               /* I:    Number of correctly sorted output positions   */
)
{
    SKP_int32    value;
    SKP_int        i, j;

    /* Safety checks */
    SKP_assert( K >  0 );
    SKP_assert( L >  0 );
    SKP_assert( L >= K );

    /* Write start indices in index vector */
    for( i = 0; i < K; i++ ) {
        index[ i ] = i;
    }

    /* Sort vector elements by value, increasing order */
    for( i = 1; i < K; i++ ) {
        value = a[ i ];
        for( j = i - 1; ( j >= 0 ) && ( value < a[ j ] ); j-- ) {
            a[ j + 1 ]     = a[ j ];     /* Shift value */
            index[ j + 1 ] = index[ j ]; /* Shift index */
        }
        a[ j + 1 ]     = value; /* Write value */
        index[ j + 1 ] = i;     /* Write index */
    }

    /* If less than L values are asked for, check the remaining values, */
    /* but only spend CPU to ensure that the K first values are correct */
    for( i = K; i < L; i++ ) {
        value = a[ i ];
        if( value < a[ K - 1 ] ) {
            for( j = K - 2; ( j >= 0 ) && ( value < a[ j ] ); j-- ) {
                a[ j + 1 ]     = a[ j ];     /* Shift value */
                index[ j + 1 ] = index[ j ]; /* Shift index */
            }
            a[ j + 1 ]     = value; /* Write value */
            index[ j + 1 ] = i;        /* Write index */
        }
    }
}

void SKP_Silk_insertion_sort_decreasing_int16(
    SKP_int16           *a,             /* I/O: Unsorted / Sorted vector                */
    SKP_int             *index,         /* O:   Index vector for the sorted elements    */
    const SKP_int       L,              /* I:   Vector length                           */
    const SKP_int       K               /* I:   Number of correctly sorted output positions    */
)
{
    SKP_int i, j;
    SKP_int value;

    /* Safety checks */
    SKP_assert( K >  0 );
    SKP_assert( L >  0 );
    SKP_assert( L >= K );

    /* Write start indices in index vector */
    for( i = 0; i < K; i++ ) {
        index[ i ] = i;
    }

    /* Sort vector elements by value, decreasing order */
    for( i = 1; i < K; i++ ) {
        value = a[ i ];
        for( j = i - 1; ( j >= 0 ) && ( value > a[ j ] ); j-- ) {    
            a[ j + 1 ]     = a[ j ];     /* Shift value */
            index[ j + 1 ] = index[ j ]; /* Shift index */
        }
        a[ j + 1 ]     = value; /* Write value */
        index[ j + 1 ] = i;     /* Write index */
    }

    /* If less than L values are asked for, check the remaining values, */
    /* but only spend CPU to ensure that the K first values are correct */
    for( i = K; i < L; i++ ) {
        value = a[ i ];
        if( value > a[ K - 1 ] ) {
            for( j = K - 2; ( j >= 0 ) && ( value > a[ j ] ); j-- ) {    
                a[ j + 1 ]     = a[ j ];     /* Shift value */
                index[ j + 1 ] = index[ j ]; /* Shift index */
            }
            a[ j + 1 ]     = value; /* Write value */
            index[ j + 1 ] = i;     /* Write index */
        }
    }
}

void SKP_Silk_insertion_sort_increasing_all_values(
    SKP_int             *a,             /* I/O: Unsorted / Sorted vector                */
    const SKP_int       L               /* I:   Vector length                           */
)
{
    SKP_int    value;
    SKP_int    i, j;

    /* Safety checks */
    SKP_assert( L >  0 );

    /* Sort vector elements by value, increasing order */
    for( i = 1; i < L; i++ ) {
        value = a[ i ];
        for( j = i - 1; ( j >= 0 ) && ( value < a[ j ] ); j-- ) {
            a[ j + 1 ] = a[ j ]; /* Shift value */
        }
        a[ j + 1 ] = value; /* Write value */
    }
}


