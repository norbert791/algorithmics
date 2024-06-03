package main

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/twmb/murmur3"
)

// text contains the first paragraph of the book "The Count of Monte Cristo" by Alexandre Dumas.
const text = `On the 24th of February, 1815, the look-out at Notre-Dame de la Garde signalled the three-master, the Pharaon from Smyrna, Trieste, and Naples.
As usual, a pilot put off immediately, and rounding the Château d’If, got on board the vessel between Cape Morgiou and Rion island.
Immediately, and according to custom, the ramparts of Fort Saint-Jean were covered with spectators; it is always an event at Marseilles for a ship to come into port, especially when this ship, like the Pharaon, has been built, rigged, and laden at the old Phocee docks, and belongs to an owner of the city.
The ship drew on and had safely passed the strait, which some volcanic shock has made between the Calasareigne and Jaros islands; had doubled Pomègue, and approached the harbor under topsails, jib, and spanker, but so slowly and sedately that the idlers, with that instinct which is the forerunner of evil, asked one another what misfortune could have happened on board. However, those experienced in navigation saw plainly that if any accident had occurred, it was not to the vessel herself, for she bore down with all the evidence of being skilfully handled, the anchor a-cockbill, the jib-boom guys already eased off, and standing by the side of the pilot, who was steering the Pharaon towards the narrow entrance of the inner port, was a young man, who, with activity and vigilant eye, watched every motion of the ship, and repeated each direction of the pilot.
The vague disquietude which prevailed among the spectators had so much affected one of the crowd that he did not await the arrival of the vessel in harbor, but jumping into a small skiff, desired to be pulled alongside the Pharaon, which he reached as she rounded into La Réserve basin.
When the young man on board saw this person approach, he left his station by the pilot, and, hat in hand, leaned over the ship’s bulwarks.
He was a fine, tall, slim young fellow of eighteen or twenty, with black eyes, and hair as dark as a raven’s wing; and his whole appearance bespoke that calmness and resolution peculiar to men accustomed from their cradle to contend with danger.
“Ah, is it you, Dantès?” cried the man in the skiff. “What’s the matter? and why have you such an air of sadness aboard?”
“A great misfortune, M. Morrel,” replied the young man, “a great misfortune, for me especially! Off Civita Vecchia we lost our brave Captain Leclere.”
“And the cargo?” inquired the owner, eagerly.
“Is all safe, M. Morrel; and I think you will be satisfied on that head. But poor Captain Leclere——”
“What happened to him?” asked the owner, with an air of considerable resignation. “What happened to the worthy captain?”
“He died.”
“Fell into the sea?”
“No, sir, he died of brain-fever in dreadful agony.” Then turning to the crew, he said, “Bear a hand there, to take in sail!”
All hands obeyed, and at once the eight or ten seamen who composed the crew, sprang to their respective stations at the spanker brails and outhaul, topsail sheets and halyards, the jib downhaul, and the topsail clewlines and buntlines. The young sailor gave a look to see that his orders were promptly and accurately obeyed, and then turned again to the owner.
`

func main() {
	seeds := []uint64{10, 20, 100, 1000, 10000}
	text := strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, "’", "")
	text = strings.ReplaceAll(text, "“", "")
	text = strings.ReplaceAll(text, "”", "")
	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, ";", "")
	text = strings.ReplaceAll(text, ":", "")
	text = strings.ReplaceAll(text, "(", "")
	text = strings.ReplaceAll(text, ")", "")
	text = strings.ReplaceAll(text, ".", "")

	splitted := strings.Split(text, " ")

	histogram := make([]uint64, 21)

	// Multiply and Add (MAD) method.
	mad := func(x uint64) uint64 {
		return ((5*x + 3) % 37) % 21
	}

	for _, seed := range seeds {
		h := murmur3.SeedNew64(seed)
		hash := func(s string) uint64 {
			_, err := h.Write([]byte(s))
			if err != nil {
				panic(err)
			}
			v := h.Sum64()
			return mad(v)
		}
		for _, word := range splitted {
			histogram[hash(word)]++
		}
	}

	fmt.Println(histogram)

	s1 := "avocado"
	s2 := "banana"

	var colCounter int

	for range 1000 {
		seed := rand.Int64()
		h := murmur3.SeedNew64(uint64(seed))
		_, err := h.Write([]byte(s1))
		if err != nil {
			panic(err)
		}
		v1 := h.Sum64()
		v1 = mad(v1)
		h.Reset()
		_, err = h.Write([]byte(s2))
		if err != nil {
			panic(err)
		}
		v2 := h.Sum64()
		v2 = mad(v2)
		if v1 == v2 {
			fmt.Println("Collision found with seed:", seed)
			colCounter++
		}
	}
	fmt.Println("Number of collisions found:", colCounter)
}
