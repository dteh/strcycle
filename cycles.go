package strcycle

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Iter is the basic struct containing the setting(s) & channel
type Iter struct {
	channel          chan interface{}
	indicateFinished bool
}

func newIter(indicateFinished bool) Iter {
	i := Iter{
		make(chan interface{}),
		indicateFinished,
	}
	return i
}

// Cycle through values forever
func Cycle(anything []interface{}, indicateFinished bool) (Iter, error) {
	if len(anything) == 0 {
		return Iter{}, errors.New("Cycle requires at least one element")
	}
	it := newIter(indicateFinished)
	go func() {
		for {
			for _, thing := range anything {
				it.channel <- thing
			}
			if indicateFinished {
				it.channel <- "**FINISHED**"
			}
		}
	}()
	return it, nil
}

// FromStringArray - Create a cycle from a string array
func FromStringArray(arr []string, indicateFinished bool) (Iter, error) {
	if len(arr) == 0 {
		return Iter{}, errors.New("FromStringArray requires at least one element")
	}
	it := newIter(indicateFinished)
	go func() {
		for {
			for _, val := range arr {
				it.channel <- val
			}
			if indicateFinished {
				it.channel <- "**FINISHED**" // indicate combinations exhausted
			}
		}
	}()
	return it, nil
}

// FromStringWithModifiers - Create cycle from a base string with str mods
func FromStringWithModifiers(base string, mods []StringModifier, indicateFinished bool) (Iter, error) {
	if len(mods) == 0 || base == "" {
		return Iter{}, errors.New("FromStringWithModifiers requires at least one StringModifier and a non empty base string")
	}
	base = strings.ToLower(base)
	it := newIter(indicateFinished)
	go func() {
		mask := make([]int, len(base))
		for i := 0; i < len(base); i++ {
			char := string(base[i])

			// Generate the mask, set bit to -1 if char is to be ignored
			if _, ok := illegalChars[char]; ok {
				mask[i] = -1
			} else {
				mask[i] = 0
			}
		}
		maskLimit := len(mods)

		for {
			mask = incrementMask(mask, maskLimit)

			builder := ""
			for i, val := range mask {
				if val == -1 {
					builder += string(base[i])
				} else {
					// builder += string(components[val][i])
					builder += string(mods[val](string(base[i])))
				}
			}
			it.channel <- builder

			// notify if mask has exhausted combos
			if indicateFinished && maskHasReset(mask) {
				it.channel <- "**FINISHED**"
				continue
			}
		}
	}()
	return it, nil
}

// ChainStringArrayWithModifiers - create an iterator cycling through an array, with string modifiers
func ChainStringArrayWithModifiers(arr []string, mods []StringModifier, indicateFinished bool) (Iter, error) {
	if len(arr) == 0 {
		return Iter{}, errors.New("ChainStringArrayWithModifiers requires at least one element")
	}
	iters := []interface{}{}
	for _, val := range arr {
		iter, err := FromStringWithModifiers(val, mods, true)
		if err != nil {
			return Iter{}, err
		}
		iters = append(iters, iter)
	}
	log.Println(iters)
	it := newIter(indicateFinished)

	go func() {
		for {
			for _, v := range iters {
				for {
					val := <-v.(Iter).channel
					if val == "**FINISHED**" {
						break
					} else {
						it.channel <- val
					}
				}
			}
			if indicateFinished {
				it.channel <- "**FINISHED**"
			}
		}
	}()
	return it, nil
}

// FormatStringWithCycles - create an iterator providing all permutations for given fstring and iterators
func FormatStringWithCycles(base string, iters ...Iter) (Iter, error) {
	if len(iters) == 0 {
		return Iter{}, errors.New("FormatStringWithCycles requires at least one iterator")
	}
	for _, iter := range iters {
		if iter.indicateFinished != true {
			return Iter{}, errors.New("All iterators must have indicateFinished=true")
		}
	}

	values := make([]interface{}, len(iters))
	it := newIter(false)
	go func() {
		// Set initial values
		for i, iter := range iters {
			if values[i] == nil {
				value := (<-iter.channel).(string)
				values[i] = value
			}
		}

		for {
			for i, iter := range iters {
				next := <-iter.channel
				if next != "**FINISHED**" {
					values[i] = next
					break
				}
				values[i] = <-iter.channel
			}
			it.channel <- fmt.Sprintf(base, values...)
		}
	}()
	return it, nil
}
