package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

/**
 * Create a window widget
 * @param parent    pointer to a parent widget
 * @return          the created window
 */
// lv_obj_t * lv_win_create(lv_obj_t * parent);

//go:linkname LvWinCreate C.lv_win_create
func LvWinCreate(parent *c.Void) *c.Void

/**
 * Add a title to the window
 * @param obj       pointer to a window widget
 * @param txt       the text of the title
 * @return          the widget where the content of the title can be created
 */
// lv_obj_t * lv_win_add_title(lv_obj_t * win, const char * txt);

//go:linkname LvWinAddTitle C.lv_win_add_title
func LvWinAddTitle(win *c.Void, txt *c.Char) *c.Void

/**
 * Add a button to the window
 * @param obj       pointer to a window widget
 * @param icon      an icon to be displayed on the button
 * @param btn_w     width of the button
 * @return          the widget where the content of the button can be created
 */
//lv_obj_t * lv_win_add_button(lv_obj_t * win, const void * icon, int32_t btn_w);

//go:linkname LvWinAddButton C.lv_win_add_button
func LvWinAddButton(win *c.Void, icon *c.Void, btn_w c.Int) *c.Void

/**
 * Get the header of the window
 * @param win       pointer to a window widget
 * @return          the header of the window
 */
// lv_obj_t * lv_win_get_header(lv_obj_t * win);

//go:linkname LvWinGetHeader C.lv_win_get_header
func LvWinGetHeader(win *c.Void) *c.Void

/**
 * Get the content of the window
 * @param win       pointer to a window widget
 * @return          the content of the window
 */
// lv_obj_t * lv_win_get_content(lv_obj_t * win);

//go:linkname LvWinGetContent C.lv_win_get_content
func LvWinGetContent(win *c.Void) *c.Void
