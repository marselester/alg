/*
Package strsort takes advantage of special properties of strings to more efficiently
sort string keys than the general-purpose sorts.

Least-significant-digit (LSD) method examines the characters in the keys in a right-to-left order.
This approach is the method of choice for string-sorting applications where all the keys are the same length.

Most-significant-digit (MSD) method examines the characters in the keys in a left-to-right order.
MSD is attractive because it can sort strings without necessarily examining all of the of the input characters.
*/
package strsort
