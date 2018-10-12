package telegram

func (ima *InputMediaAnimation) File() string                { return ima.Media }
func (ima *InputMediaAnimation) InputMediaCaption() string   { return ima.Caption }
func (ima *InputMediaAnimation) InputMediaParseMode() string { return ima.ParseMode }
func (ima *InputMediaAnimation) InputMediaType() string      { return ima.Type }

func (imd *InputMediaDocument) File() string                { return imd.Media }
func (imd *InputMediaDocument) InputMediaCaption() string   { return imd.Caption }
func (imd *InputMediaDocument) InputMediaParseMode() string { return imd.ParseMode }
func (imd *InputMediaDocument) InputMediaType() string      { return imd.Type }

func (ima *InputMediaAudio) File() string                { return ima.Media }
func (ima *InputMediaAudio) InputMediaCaption() string   { return ima.Caption }
func (ima *InputMediaAudio) InputMediaParseMode() string { return ima.ParseMode }
func (ima *InputMediaAudio) InputMediaType() string      { return ima.Type }

func (imp *InputMediaPhoto) File() string                { return imp.Media }
func (imp *InputMediaPhoto) InputMediaCaption() string   { return imp.Caption }
func (imp *InputMediaPhoto) InputMediaParseMode() string { return imp.ParseMode }
func (imp *InputMediaPhoto) InputMediaType() string      { return imp.Type }

func (imv *InputMediaVideo) File() string                { return imv.Media }
func (imv *InputMediaVideo) InputMediaCaption() string   { return imv.Caption }
func (imv *InputMediaVideo) InputMediaParseMode() string { return imv.ParseMode }
func (imv *InputMediaVideo) InputMediaType() string      { return imv.Type }
