// *
// * Copyright 2014, Scott Cagno, All rights reserved.
// * BSD Licensed - sites.google.com/site/bsdc3license
// *
// * B+Tree Implementation
// *
package tree

// basic node `rules`, i think
//
// root --> ((1 | 2 ) <= children <= deg)
// node --> ((deg/2) <= children <= deg)
// leaf --> ((deg/2) <= records <= (deg-1))
//
// deg := 6
// root has between 1 or 2 and 6 children
// node has between 3 and 6 children
// leaf has between 3 and 5 records
