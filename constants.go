// This file is automatically generated. Do not edit.

package pyx

import ()

// AjaxOperation
const (
	AjaxOperation_NAMES = "gn"
	AjaxOperation_SCORE = "SC"
	AjaxOperation_LEAVE_GAME = "lg"
	AjaxOperation_JOIN_GAME = "jg"
	AjaxOperation_CHAT = "c"
	AjaxOperation_GAME_LIST = "ggl"
	AjaxOperation_CARDCAST_ADD_CARDSET = "cac"
	AjaxOperation_CARDCAST_LIST_CARDSETS = "clc"
	AjaxOperation_PLAY_CARD = "pc"
	AjaxOperation_CHANGE_GAME_OPTIONS = "cgo"
	AjaxOperation_GET_GAME_INFO = "ggi"
	AjaxOperation_GET_CARDS = "gc"
	AjaxOperation_ADMIN_SET_VERBOSE_LOG = "svl"
	AjaxOperation_REGISTER = "r"
	AjaxOperation_CARDCAST_REMOVE_CARDSET = "crc"
	AjaxOperation_WHOIS = "Wi"
	AjaxOperation_KICK = "K"
	AjaxOperation_FIRST_LOAD = "fl"
	AjaxOperation_START_GAME = "sg"
	AjaxOperation_LOG_OUT = "lo"
	AjaxOperation_BAN = "b"
	AjaxOperation_CREATE_GAME = "cg"
	AjaxOperation_STOP_GAME = "Sg"
	AjaxOperation_GAME_CHAT = "GC"
	AjaxOperation_SPECTATE_GAME = "vg"
	AjaxOperation_JUDGE_SELECT = "js"
)


// AjaxRequest
const (
	AjaxRequest_SERIAL = "s"
	AjaxRequest_OP = "o"
	AjaxRequest_WALL = "wall"
	AjaxRequest_PERSISTENT_ID = "pid"
	AjaxRequest_EMOTE = "me"
	AjaxRequest_CARDCAST_ID = "cci"
	AjaxRequest_GAME_ID = "gid"
	AjaxRequest_GAME_OPTIONS = "go"
	AjaxRequest_MESSAGE = "m"
	AjaxRequest_NICKNAME = "n"
	AjaxRequest_PASSWORD = "pw"
	AjaxRequest_CARD_ID = "cid"
	AjaxRequest_ID_CODE = "idc"
)

type AjaxRequest struct {
	Serial int `json:"s"`
	Op string `json:"o"`
	Wall bool `json:"wall"`
	PersistentId string `json:"pid"`
	Emote bool `json:"me"`
	CardcastId string `json:"cci"`
	GameId int `json:"gid"`
	GameOptions GameOptionData `json:"go"`
	Message string `json:"m"`
	Nickname string `json:"n"`
	Password string `json:"pw"`
	CardId int `json:"cid"`
	IdCode string `json:"idc"`
}


// AjaxResponse
const (
	AjaxResponse_NAMES = "nl"
	AjaxResponse_CLIENT_NAME = "cn"
	AjaxResponse_GAME_CHAT_ENABLED = "Gce"
	AjaxResponse_PLAYER_INFO = "pi"
	AjaxResponse_CONNECTED_AT = "ca"
	AjaxResponse_WHITE_CARDS = "wc"
	AjaxResponse_HAND = "h"
	AjaxResponse_ERROR_CODE = "ec"
	AjaxResponse_GLOBAL_CHAT_ENABLED = "gce"
	AjaxResponse_SERVER_STARTED = "SS"
	AjaxResponse_NEXT = "next"
	AjaxResponse_GAME_INFO = "gi"
	AjaxResponse_ERROR = "e"
	AjaxResponse_GAME_STATE_DESCRIPTION = "gss"
	AjaxResponse_ID_CODE = "idc"
	AjaxResponse_CARD_SETS = "css"
	AjaxResponse_SERIAL = "s"
	AjaxResponse_PERSISTENT_ID = "pid"
	AjaxResponse_USER_PERMALINK = "up"
	AjaxResponse_GAMES = "gl"
	AjaxResponse_SIGIL = "?"
	AjaxResponse_GAME_ID = "gid"
	AjaxResponse_MAX_GAMES = "mg"
	AjaxResponse_IN_PROGRESS = "ip"
	AjaxResponse_GAME_OPTIONS = "go"
	AjaxResponse_NICKNAME = "n"
	AjaxResponse_BLACK_CARD = "bc"
	AjaxResponse_GAME_PERMALINK = "gp"
	AjaxResponse_IDLE = "idl"
	AjaxResponse_CARD_ID = "cid"
	AjaxResponse_IP_ADDRESS = "IP"
	AjaxResponse_SESSION_PERMALINK = "sP"
)

type AjaxResponse struct {
	Names []string `json:"nl"`
	ClientName string `json:"cn"`
	GameChatEnabled bool `json:"Gce"`
	PlayerInfo []GamePlayerInfo `json:"pi"`
	ConnectedAt int64 `json:"ca"`
	WhiteCards []int `json:"wc"`
	Hand []int `json:"h"`
	ErrorCode string `json:"ec"`
	GlobalChatEnabled bool `json:"gce"`
	ServerStarted int64 `json:"SS"`
	Next string `json:"next"`
	GameInfo GameInfo `json:"gi"`
	Error bool `json:"e"`
	GameStateDescription string `json:"gss"`
	IdCode string `json:"idc"`
	CardSets []CardSetData `json:"css"`
	Serial int `json:"s"`
	PersistentId string `json:"pid"`
	UserPermalink string `json:"up"`
	Games []GameInfo `json:"gl"`
	Sigil string `json:"?"`
	GameId *int `json:"gid"`
	MaxGames int `json:"mg"`
	InProgress bool `json:"ip"`
	GameOptions GameOptionData `json:"go"`
	Nickname string `json:"n"`
	BlackCard int `json:"bc"`
	GamePermalink string `json:"gp"`
	Idle int64 `json:"idl"`
	CardId int `json:"cid"`
	IpAddress string `json:"IP"`
	SessionPermalink string `json:"sP"`
}


// BlackCardData
const (
	BlackCardData_DRAW = "D"
	BlackCardData_PICK = "PK"
	BlackCardData_TEXT = "T"
	BlackCardData_ID = "cid"
	BlackCardData_WATERMARK = "W"
)

type BlackCardData struct {
	Draw int `json:"D"`
	Pick int `json:"PK"`
	Text string `json:"T"`
	Id int `json:"cid"`
	Watermark string `json:"W"`
}


// CardSetData
const (
	CardSetData_WHITE_CARDS_IN_DECK = "wcid"
	CardSetData_BLACK_CARDS_IN_DECK = "bcid"
	CardSetData_CARD_SET_NAME = "csn"
	CardSetData_CARD_SET_DESCRIPTION = "csd"
	CardSetData_BASE_DECK = "bd"
	CardSetData_ID = "cid"
	CardSetData_WEIGHT = "w"
)

type CardSetData struct {
	WhiteCardsInDeck int `json:"wcid"`
	BlackCardsInDeck int `json:"bcid"`
	CardSetName string `json:"csn"`
	CardSetDescription string `json:"csd"`
	BaseDeck bool `json:"bd"`
	Id int `json:"cid"`
	Weight int `json:"w"`
}


// DisconnectReason
const (
	DisconnectReason_PING_TIMEOUT = "pt"
	DisconnectReason_BANNED = "B&"
	DisconnectReason_IDLE_TIMEOUT = "it"
	DisconnectReason_KICKED = "k"
	DisconnectReason_MANUAL = "man"
)

var DisconnectReasonMsgs = map[string]string{
	 "B&": "Banned",
	 "pt": "Ping timeout",
	 "it": "Kicked due to idle",
	 "k": "Kicked by server administrator",
	 "man": "Leaving",
}


// ErrorCode
const (
	ErrorCode_INVALID_NICK = "in"
	ErrorCode_TOO_MANY_USERS = "tmu"
	ErrorCode_NOT_GAME_HOST = "ngh"
	ErrorCode_MESSAGE_TOO_LONG = "mtl"
	ErrorCode_NO_NICK_SPECIFIED = "nns"
	ErrorCode_REPEATED_WORDS = "rW"
	ErrorCode_BANNED = "B&"
	ErrorCode_WRONG_PASSWORD = "wp"
	ErrorCode_RESERVED_NICK = "rn"
	ErrorCode_TOO_MANY_GAMES = "tmg"
	ErrorCode_INVALID_ID_CODE = "iid"
	ErrorCode_CANNOT_JOIN_ANOTHER_GAME = "cjag"
	ErrorCode_NO_MSG_SPECIFIED = "nms"
	ErrorCode_CAPSLOCK = "CL"
	ErrorCode_ALREADY_STARTED = "as"
	ErrorCode_NOT_ADMIN = "na"
	ErrorCode_INVALID_GAME = "ig"
	ErrorCode_NO_SESSION = "ns"
	ErrorCode_ACCESS_DENIED = "ad"
	ErrorCode_NICK_IN_USE = "niu"
	ErrorCode_NOT_JUDGE = "nj"
	ErrorCode_SERVER_ERROR = "serr"
	ErrorCode_CARDCAST_INVALID_ID = "cii"
	ErrorCode_TOO_FAST = "tf"
	ErrorCode_NOT_ENOUGH_CARDS = "nec"
	ErrorCode_NO_CARD_SPECIFIED = "ncs"
	ErrorCode_REPEAT_MESSAGE = "rm"
	ErrorCode_NO_GAME_SPECIFIED = "ngs"
	ErrorCode_OP_NOT_SPECIFIED = "ons"
	ErrorCode_TOO_MANY_SPECIAL_CHARACTERS = "tmsc"
	ErrorCode_BAD_REQUEST = "br"
	ErrorCode_NOT_ENOUGH_PLAYERS = "nep"
	ErrorCode_CARDCAST_CANNOT_FIND = "ccf"
	ErrorCode_NOT_IN_THAT_GAME = "nitg"
	ErrorCode_NO_SUCH_USER = "nsu"
	ErrorCode_NOT_REGISTERED = "nr"
	ErrorCode_BAD_OP = "bo"
	ErrorCode_DO_NOT_HAVE_CARD = "dnhc"
	ErrorCode_NOT_YOUR_TURN = "nyt"
	ErrorCode_PLAYED_ALL_CARDS = "pac"
	ErrorCode_NOT_ENOUGH_SPACES = "nes"
	ErrorCode_ALREADY_STOPPED = "aS"
	ErrorCode_SESSION_EXPIRED = "se"
	ErrorCode_GAME_FULL = "gf"
	ErrorCode_INVALID_CARD = "ic"
)

var ErrorCodeMsgs = map[string]string{
	 "cii": "Invalid Cardcast ID. Must be exactly 5 characters.",
	 "nr": "Not registered. Refresh the page.",
	 "iid": "Identification code, if provided, must be between 8 and 100 characters, inclusive.",
	 "ns": "Session not detected. Make sure you have cookies enabled.",
	 "ccf": "Cannot find Cardcast deck with given ID. If you just added this deck to Cardcast, wait a few minutes and try again.",
	 "nyt": "It is not your turn to play a card.",
	 "bo": "Invalid operation.",
	 "nec": "You must add card sets containing at least 50 black cards and 20 times the player limit white cards.",
	 "ngh": "Only the game host can do that.",
	 "tmg": "There are too many games already in progress. Either join an existing game, or wait for one to become available.",
	 "br": "Bad request.",
	 "nsu": "No such user.",
	 "aS": "The game has already stopped.",
	 "se": "Your session has expired. Refresh the page.",
	 "pac": "You already played all the necessary cards!",
	 "nms": "No message specified.",
	 "nep": "There are not enough players to start the game.",
	 "wp": "That password is incorrect.",
	 "ic": "Invalid card specified.",
	 "niu": "Nickname is already in use.",
	 "ngs": "No game specified.",
	 "nes": "You must use more words in a message that long.",
	 "nitg": "You are not in that game.",
	 "tmu": "There are too many users connected. Either join another server, or wait for a user to disconnect.",
	 "ig": "Invalid game specified.",
	 "gf": "That game is full. Join another.",
	 "ncs": "No card specified.",
	 "ad": "Access denied.",
	 "cjag": "You cannot join another game.",
	 "B&": "Banned.",
	 "mtl": "Messages cannot be longer than 200 characters.",
	 "in": "Nickname must contain only upper and lower case letters, numbers, or underscores, must be 3 to 30 characters long, and must not start with a number.",
	 "rW": "You must use more unique words in your message.",
	 "serr": "An error occurred on the server.",
	 "CL": "Try turning caps lock off.",
	 "dnhc": "You don't have that card.",
	 "as": "The game has already started.",
	 "nns": "No nickname specified.",
	 "tf": "You are chatting too fast. Wait a few seconds and try again.",
	 "na": "You are not an administrator.",
	 "ons": "Operation not specified.",
	 "rm": "You can't repeat the same message multiple times in a row.",
	 "nj": "You are not the judge.",
	 "rn": "That nick is reserved.",
	 "tmsc": "You used too many special characters in that message.",
}


// ErrorInformation
const (
	ErrorInformation_WHITE_CARDS_PRESENT = "wcp"
	ErrorInformation_WHITE_CARDS_REQUIRED = "wcr"
	ErrorInformation_BLACK_CARDS_REQUIRED = "bcr"
	ErrorInformation_BLACK_CARDS_PRESENT = "bcp"
)


// GameInfo
const (
	GameInfo_GAME_OPTIONS = "go"
	GameInfo_CREATED = "gca"
	GameInfo_PLAYERS = "P"
	GameInfo_SPECTATORS = "V"
	GameInfo_HOST = "H"
	GameInfo_STATE = "S"
	GameInfo_ID = "gid"
	GameInfo_HAS_PASSWORD = "hp"
)

type GameInfo struct {
	GameOptions GameOptionData `json:"go"`
	Created int64 `json:"gca"`
	Players []string `json:"P"`
	Spectators []string `json:"V"`
	Host string `json:"H"`
	State string `json:"S"`
	Id int `json:"gid"`
	HasPassword bool `json:"hp"`
}


// GameOptionData
const (
	GameOptionData_TIMER_MULTIPLIER = "tm"
	GameOptionData_PASSWORD = "pw"
	GameOptionData_SPECTATOR_LIMIT = "vL"
	GameOptionData_SCORE_LIMIT = "sl"
	GameOptionData_BLANKS_LIMIT = "bl"
	GameOptionData_PLAYER_LIMIT = "pL"
	GameOptionData_CARD_SETS = "css"
)

type GameOptionData struct {
	TimerMultiplier string `json:"tm"`
	Password string `json:"pw"`
	SpectatorLimit int `json:"vL"`
	ScoreLimit int `json:"sl"`
	BlanksLimit int `json:"bl"`
	PlayerLimit int `json:"pL"`
	CardSets []int `json:"css"`
}


// GamePlayerInfo
const (
	GamePlayerInfo_STATUS = "st"
	GamePlayerInfo_SCORE = "sc"
	GamePlayerInfo_NAME = "N"
)

type GamePlayerInfo struct {
	Status string `json:"st"`
	Score int `json:"sc"`
	Name string `json:"N"`
}


// GamePlayerStatus
const (
	GamePlayerStatus_SPECTATOR = "sv"
	GamePlayerStatus_WINNER = "sw"
	GamePlayerStatus_IDLE = "si"
	GamePlayerStatus_HOST = "sh"
	GamePlayerStatus_JUDGING = "sjj"
	GamePlayerStatus_JUDGE = "sj"
	GamePlayerStatus_PLAYING = "sp"
)

var GamePlayerStatusMsgs = map[string]string{
	 "sjj": "Selecting",
	 "sv": "Spectator",
	 "sw": "Winner!",
	 "sh": "Host",
	 "si": "",
	 "sj": "Card Czar",
	 "sp": "Playing",
}

var GamePlayerStatusMsgs2 = map[string]string{
	 "sjj": "Select a winning card.",
	 "sv": "You are just spectating.",
	 "sw": "You have won!",
	 "sh": "Wait for players then click Start Game.",
	 "si": "Waiting for players...",
	 "sj": "You are the Card Czar.",
	 "sp": "Select a card to play.",
}


// GameState
const (
	GameState_ROUND_OVER = "ro"
	GameState_LOBBY = "l"
	GameState_JUDGING = "j"
	GameState_PLAYING = "p"
)

var GameStateMsgs = map[string]string{
	 "p": "In Progress",
	 "j": "In Progress",
	 "l": "Not Started",
	 "ro": "In Progress",
}


// LongPollEvent
const (
	LongPollEvent_GAME_ROUND_COMPLETE = "grc"
	LongPollEvent_BANNED = "B&"
	LongPollEvent_NOOP = "_"
	LongPollEvent_CHAT = "c"
	LongPollEvent_GAME_PLAYER_INFO_CHANGE = "gpic"
	LongPollEvent_HAND_DEAL = "hd"
	LongPollEvent_CARDCAST_ADD_CARDSET = "cac"
	LongPollEvent_PLAYER_LEAVE = "pl"
	LongPollEvent_GAME_BLACK_RESHUFFLE = "gbr"
	LongPollEvent_GAME_JUDGE_SKIPPED = "gjs"
	LongPollEvent_GAME_LIST_REFRESH = "glr"
	LongPollEvent_NEW_PLAYER = "np"
	LongPollEvent_GAME_PLAYER_SKIPPED = "gps"
	LongPollEvent_GAME_PLAYER_JOIN = "gpj"
	LongPollEvent_GAME_WHITE_RESHUFFLE = "gwr"
	LongPollEvent_CARDCAST_REMOVE_CARDSET = "crc"
	LongPollEvent_GAME_OPTIONS_CHANGED = "goc"
	LongPollEvent_GAME_PLAYER_KICKED_IDLE = "gpki"
	LongPollEvent_GAME_SPECTATOR_LEAVE = "gvl"
	LongPollEvent_GAME_PLAYER_LEAVE = "gpl"
	LongPollEvent_GAME_SPECTATOR_JOIN = "gvj"
	LongPollEvent_HURRY_UP = "hu"
	LongPollEvent_GAME_JUDGE_LEFT = "gjl"
	LongPollEvent_KICKED = "k"
	LongPollEvent_KICKED_FROM_GAME_IDLE = "kfgi"
	LongPollEvent_FILTERED_CHAT = "FC"
	LongPollEvent_GAME_STATE_CHANGE = "gsc"
)


// LongPollResponse
const (
	LongPollResponse_PLAY_TIMER = "Pt"
	LongPollResponse_PLAYER_INFO = "pi"
	LongPollResponse_FROM = "f"
	LongPollResponse_WHITE_CARDS = "wc"
	LongPollResponse_EVENT = "E"
	LongPollResponse_HAND = "h"
	LongPollResponse_ERROR_CODE = "ec"
	LongPollResponse_MESSAGE = "m"
	LongPollResponse_WINNING_CARD = "WC"
	LongPollResponse_FROM_ADMIN = "fa"
	LongPollResponse_TIMESTAMP = "ts"
	LongPollResponse_GAME_INFO = "gi"
	LongPollResponse_ERROR = "e"
	LongPollResponse_ID_CODE = "idc"
	LongPollResponse_REASON = "qr"
	LongPollResponse_WALL = "wall"
	LongPollResponse_ROUND_WINNER = "rw"
	LongPollResponse_SIGIL = "?"
	LongPollResponse_EMOTE = "me"
	LongPollResponse_CARDCAST_DECK_INFO = "cdi"
	LongPollResponse_GAME_ID = "gid"
	LongPollResponse_ROUND_PERMALINK = "rP"
	LongPollResponse_NICKNAME = "n"
	LongPollResponse_BLACK_CARD = "bc"
	LongPollResponse_GAME_PERMALINK = "gp"
	LongPollResponse_GAME_STATE = "gs"
	LongPollResponse_INTERMISSION = "i"
)

type LongPollResponse struct {
	PlayTimer int `json:"Pt"`
	PlayerInfo []GamePlayerInfo `json:"pi"`
	From string `json:"f"`
	WhiteCards [][]WhiteCardData `json:"wc"`
	Event string `json:"E"`
	Hand []WhiteCardData `json:"h"`
	ErrorCode string `json:"ec"`
	Message string `json:"m"`
	WinningCard int `json:"WC"`
	FromAdmin bool `json:"fa"`
	Timestamp int64 `json:"ts"`
	GameInfo GameInfo `json:"gi"`
	Error bool `json:"e"`
	IdCode string `json:"idc"`
	Reason string `json:"qr"`
	Wall bool `json:"wall"`
	RoundWinner string `json:"rw"`
	Sigil string `json:"?"`
	Emote bool `json:"me"`
	CardcastDeckInfo string `json:"cdi"`
	GameId *int `json:"gid"`
	RoundPermalink string `json:"rP"`
	Nickname string `json:"n"`
	BlackCard BlackCardData `json:"bc"`
	GamePermalink string `json:"gp"`
	GameState string `json:"gs"`
	Intermission int `json:"i"`
}


// ReconnectNextAction
const (
	ReconnectNextAction_GAME = "game"
	ReconnectNextAction_NONE = "none"
)


// Sigil
const (
	Sigil_NORMAL_USER = ""
	Sigil_ADMIN = "@"
	Sigil_ID_CODE = "+"
)


// WhiteCardData
const (
	WhiteCardData_WRITE_IN = "wi"
	WhiteCardData_TEXT = "T"
	WhiteCardData_ID = "cid"
	WhiteCardData_WATERMARK = "W"
)

type WhiteCardData struct {
	WriteIn bool `json:"wi"`
	Text string `json:"T"`
	Id int `json:"cid"`
	Watermark string `json:"W"`
}

