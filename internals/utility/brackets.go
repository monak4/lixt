package utility

type BracketHandler struct {
	pairs map[rune]rune
}

func NewBracketsHandler() *BracketHandler {
	return &BracketHandler{
		pairs: map[rune]rune{
			'(':  ')',
			'[':  ']',
			'{':  '}',
			'"':  '"',
			'\'': '\'',
		},
	}
}

func (h *BracketHandler) HandleBracket(text string, caretPos int, char rune) (newText string, newCaretPos int, handled bool) {
	if closing, ok := h.pairs[char]; ok {
		newText = text[:caretPos] + string(char) + string(closing) + text[caretPos:]
		newCaretPos = caretPos + 1
		return newText, newCaretPos, false
	}
	return text, caretPos, false
}

func (h *BracketHandler) ShouldSkipClosing(text string, caretPos int, char rune) bool {
	if caretPos >= len(text) {
		return false
	}
	return text[caretPos] == uint8(char)
}

func (h *BracketHandler) HandleDelete(text string, caretPos int) (newText string, newCaretPos int, handled bool) {
	if caretPos <= 0 || caretPos >= len(text) {
		return text, caretPos, false
	}

	prevChar := rune(text[caretPos-1])
	if closing, ok := h.pairs[prevChar]; ok {
		if caretPos < len(text) && rune(text[caretPos]) == closing {
			newText = text[:caretPos-1] + text[caretPos+1:]
			newCaretPos = caretPos - 1
			return newText, newCaretPos, true
		}
	}
	return text, caretPos, false
}

func (h *BracketHandler) WrapSelection(text string, selStart, selEnd int, openBracket rune) (string, int) {
	if closing, ok := h.pairs[openBracket]; ok {
		newText := text[:selStart] +
			string(openBracket) +
			text[selStart:selEnd] +
			string(closing) +
			text[selEnd:]
		newCaretPos := selEnd + 2
		return newText, newCaretPos
	}
	return text, selEnd
}
