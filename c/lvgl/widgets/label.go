package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

/**
 * Create a label object
 * @param parent    pointer to an object, it will be the parent of the new label.
 * @return          pointer to the created button
 */
//lv_obj_t * lv_label_create(lv_obj_t * parent);
//
//go:linkname LvLabelCreate C.lv_label_create
func LvLabelCreate(parent *c.Void) *c.Void

/*=====================
 * Setter functions
 *====================*/

/**
 * Set a new text for a label. Memory will be allocated to store the text by the label.
 * @param obj           pointer to a label object
 * @param text          '\0' terminated character string. NULL to refresh with the current text.
 */
//void lv_label_set_text(lv_obj_t * obj, const char * text);
//
//go:linkname LvLabelSetText C.lv_label_set_text
func LvLabelSetText(obj *c.Void, text *c.Char)

/**
 * Set a new formatted text for a label. Memory will be allocated to store the text by the label.
 * @param obj           pointer to a label object
 * @param fmt           `printf`-like format
 *
 * Example:
 * @code
 * lv_label_set_text_fmt(label1, "%d user", user_num);
 * @endcode
 */
//void lv_label_set_text_fmt(lv_obj_t * obj, const char * fmt, ...) LV_FORMAT_ATTRIBUTE(2, 3);
//
//go:linkname LvLabelSetTextFmt C.lv_label_set_text_fmt
func LvLabelSetTextFmt(obj *c.Void, fmt *c.Char, args ...interface{})

/**
 * Set a static text. It will not be saved by the label so the 'text' variable
 * has to be 'alive' while the label exists.
 * @param obj           pointer to a label object
 * @param text          pointer to a text. NULL to refresh with the current text.
 */
//void lv_label_set_text_static(lv_obj_t * obj, const char * text);
//
//go:linkname LvLabelSetTextStatic C.lv_label_set_text_static
func LvLabelSetTextStatic(obj *c.Void, text *c.Char)

/**
 * Set the behavior of the label with text longer than the object size
 * @param obj           pointer to a label object
 * @param long_mode     the new mode from 'lv_label_long_mode' enum.
 *                      In LV_LONG_WRAP/DOT/SCROLL/SCROLL_CIRC the size of the label should be set AFTER this function
 */
//void lv_label_set_long_mode(lv_obj_t * obj, lv_label_long_mode_t long_mode);
//
//go:linkname LvLabelSetLongMode C.lv_label_set_long_mode
func LvLabelSetLongMode(obj *c.Void, long_mode c.Int)

/**
 * Set where text selection should start
 * @param obj       pointer to a label object
 * @param index     character index from where selection should start. `LV_LABEL_TEXT_SELECTION_OFF` for no selection
 */
//void lv_label_set_text_selection_start(lv_obj_t * obj, uint32_t index);
//
//go:linkname LvLabelSetTextSelectionStart C.lv_label_set_text_selection_start
func LvLabelSetTextSelectionStart(obj *c.Void, index c.Uint)

/**
 * Set where text selection should end
 * @param obj       pointer to a label object
 * @param index     character index where selection should end. `LV_LABEL_TEXT_SELECTION_OFF` for no selection
 */
//void lv_label_set_text_selection_end(lv_obj_t * obj, uint32_t index);
//
//go:linkname LvLabelSetTextSelectionEnd C.lv_label_set_text_selection_end
func LvLabelSetTextSelectionEnd(obj *c.Void, index c.Uint)

/**
 * Enable the recoloring by in-line commands
 * @param obj           pointer to a label object
 * @param en            true: enable recoloring, false: disable
 * Example: "This is a #ff0000 red# word"
 */
//void lv_label_set_recolor(lv_obj_t * obj, bool en);
//
//go:linkname LvLabelSetRecolor C.lv_label_set_recolor
func LvLabelSetRecolor(obj *c.Void, en c.Char)

/*=====================
 * Getter functions
 *====================*/

/**
 * Get the text of a label
 * @param obj       pointer to a label object
 * @return          the text of the label
 */
//char * lv_label_get_text(const lv_obj_t * obj);
//
//go:linkname LvLabelGetText C.lv_label_get_text
func LvLabelGetText(obj *c.Void) *c.Char

/**
 * Get the long mode of a label
 * @param obj       pointer to a label object
 * @return          the current long mode
 */
//lv_label_long_mode_t lv_label_get_long_mode(const lv_obj_t * obj);
//
//go:linkname LvLabelGetLongMode C.lv_label_get_long_mode
func LvLabelGetLongMode(obj *c.Void) c.Int

/**
 * Get the relative x and y coordinates of a letter
 * @param obj       pointer to a label object
 * @param char_id   index of the character [0 ... text length - 1].
 *                  Expressed in character index, not byte index (different in UTF-8)
 * @param pos       store the result here (E.g. index = 0 gives 0;0 coordinates if the text if aligned to the left)
 */
//void lv_label_get_letter_pos(const lv_obj_t * obj, uint32_t char_id, lv_point_t * pos);
//
//go:linkname LvLabelGetLetterPos C.lv_label_get_letter_pos
func LvLabelGetLetterPos(obj *c.Void, char_id c.Uint, pos *c.Char)

/**
 * Get the index of letter on a relative point of a label.
 * @param obj       pointer to label object
 * @param pos_in    pointer to point with coordinates on a the label
 * @param bidi      whether to use bidi processed
 * @return          The index of the letter on the 'pos_p' point (E.g. on 0;0 is the 0. letter if aligned to the left)
 *                  Expressed in character index and not byte index (different in UTF-8)
 */
//uint32_t lv_label_get_letter_on(const lv_obj_t * obj, lv_point_t * pos_in, bool bidi);
//
//go:linkname LvLabelGetLetterOn C.lv_label_get_letter_on
func LvLabelGetLetterOn(obj *c.Void, pos_in *c.Char, bidi c.Char) c.Uint

/**
 * Check if a character is drawn under a point.
 * @param obj       pointer to a label object
 * @param pos       Point to check for character under
 * @return          whether a character is drawn under the point
 */
//bool lv_label_is_char_under_pos(const lv_obj_t * obj, lv_point_t * pos);
//
//go:linkname LvLabelIsCharUnderPos C.lv_label_is_char_under_pos
func LvLabelIsCharUnderPos(obj *c.Void, pos *c.Char) c.Char

/**
 * @brief Get the selection start index.
 * @param obj       pointer to a label object.
 * @return          selection start index. `LV_LABEL_TEXT_SELECTION_OFF` if nothing is selected.
 */
//uint32_t lv_label_get_text_selection_start(const lv_obj_t * obj);
//
//go:linkname LvLabelGetTextSelectionStart C.lv_label_get_text_selection_start
func LvLabelGetTextSelectionStart(obj *c.Void) c.Uint

/**
 * @brief Get the selection end index.
 * @param obj       pointer to a label object.
 * @return          selection end index. `LV_LABEL_TXT_SEL_OFF` if nothing is selected.
 */
//uint32_t lv_label_get_text_selection_end(const lv_obj_t * obj);
//
//go:linkname LvLabelGetTextSelectionEnd C.lv_label_get_text_selection_end
func LvLabelGetTextSelectionEnd(obj *c.Void) c.Uint

/**
 * @brief Get the recoloring attribute
 * @param obj       pointer to a label object.
 * @return          true: recoloring is enabled, false: recoloring is disabled
 */
//bool lv_label_get_recolor(const lv_obj_t * obj);
//
//go:linkname LvLabelGetRecolor C.lv_label_get_recolor
func LvLabelGetRecolor(obj *c.Void) c.Char

/*=====================
 * Other functions
 *====================*/

/**
 * Insert a text to a label. The label text cannot be static.
 * @param obj       pointer to a label object
 * @param pos       character index to insert. Expressed in character index and not byte index.
 *                  0: before first char. LV_LABEL_POS_LAST: after last char.
 * @param txt       pointer to the text to insert
 */
//void lv_label_ins_text(lv_obj_t * obj, uint32_t pos, const char * txt);
//
//go:linkname LvLabelInsText C.lv_label_ins_text
func LvLabelInsText(obj *c.Void, pos c.Uint, txt *c.Char)

/**
 * Delete characters from a label. The label text cannot be static.
 * @param obj       pointer to a label object
 * @param pos       character index from where to cut. Expressed in character index and not byte index.
 *                  0: start in front of the first character
 * @param cnt       number of characters to cut
 */
//void lv_label_cut_text(lv_obj_t * obj, uint32_t pos, uint32_t cnt);
//
//go:linkname LvLabelCutText C.lv_label_cut_text
func LvLabelCutText(obj *c.Void, pos c.Uint, cnt c.Uint)
