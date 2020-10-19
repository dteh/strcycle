# strcycle
basic tools to cycle through permutations of strings & arrays

## how2use
*all cycles take an arg `indicateFinished` which returns a string `**FINISHED**` as an element when the cycle has exhausted all permutations, and is about to restart*

- Cycle - cycle through an array of values
- FromStringArray - create cycle from an array of strings (probably should be removed)
- FromStringWithModifiers - for a given string, cycle through all permutations of string modified versions of the given string ( see string modifiers below )
    - eg. for the modifiers UPPER and LOWER the string "hey" returns "heY", "hEy", "Hey", "HEy" etc.
- FormatStringWithCycles - for a given f string, format it with all permutations of multiple cycles

## string modifiers
*type (func(s string) string)*

a function taking a string and performing some op on it then returning said string

predefined string modifiers

- UPPER - uppercase the entire string
- LOWER - lowercase the entire string
- SingleEncodeUpper - upper case the string then % encode each char
- SingleEncodeLower - lower case the string then % encode each char
- DoubleEncodeUpper - upper case the string then double % encode each char
- DoubleEncodeLower - lower case the string then double % encode each char