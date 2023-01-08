// the goal is to find the area of highlighted text in a pdf viewer
// designerPdfViewer has the following parameter(s):

//    int h[26]: the heights of letter highlight
//    string word: a string

// Caveat: The highlight's effective height is the same as the highest letter height

// Returns

//    int: the size of the highlighted area

package misc

func designerPdfViewer(h []int32, word string) int32 {
	var max_height int32 = -1
	var area int32

	for _, letter := range word {
		letter_num := letter - 'a' // assuming lowercase
		letter_height := h[letter_num]
		if letter_height > max_height {
			max_height = letter_height
		}
	}

	area = int32(len(word)) * max_height

	return area
}
