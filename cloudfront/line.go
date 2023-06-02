package cloudfront

// Line is a struct that represents a single line in a Cloudfront log file.
// @see https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html
const (
	Date = iota
	Time
	EdegeLocation
	SCBytes
	IP
	Method
	Host
	URI
	Status
	Referer
	UserAgent
	Query
	Cookie
	EdgeDefaultResultType
	EdgeRequestId
	HostHeader
	Protocol
	CSByes
	TimeTaken
	ForwardedFor
	SSLProtocol
	SSLCypher
	EdgeResponseResultType
	ProtocolVersion
	FLEStatus
	FLEEncryptedFields
	Port
	TTFB
	EdgeDetailedResultType
	ContentType
	ContentLen
	RangeStart
	RangeEnd
)

// Line is a struct that represents a single line in a CloudFront log file.
// @see https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html
type Line struct {
	Date                   string `csv:"date"`
	Time                   string `csv:"time"`
	EdegeLocation          string `csv:"x-edge-location"`
	SCBytes                string `csv:"sc-bytes"`
	IP                     string `csv:"c-ip"`
	Method                 string `csv:"cs-method"`
	Host                   string `csv:"cs(Host)"`
	URI                    string `csv:"cs-uri-stem"`
	Status                 string `csv:"sc-status"`
	Referer                string `csv:"cs(Referer)"`
	UserAgent              string `csv:"cs(User-Agent)"`
	Query                  string `csv:"cs-uri-query"`
	Cookie                 string `csv:"cs(Cookie)"`
	EdgeDefaultResultType  string `csv:"x-edge-result-type"`
	EdgeRequestId          string `csv:"x-edge-request-id"`
	HostHeader             string `csv:"x-host-header"`
	Protocol               string `csv:"cs-protocol"`
	CSByes                 string `csv:"cs-bytes"`
	TimeTaken              string `csv:"time-taken"`
	ForwardedFor           string `csv:"x-forwarded-for"`
	SSLProtocol            string `csv:"ssl-protocol"`
	SSLCypher              string `csv:"ssl-cipher"`
	EdgeResponseResultType string `csv:"x-edge-response-result-type"`
	ProtocolVersion        string `csv:"cs-protocol-version"`
	FLEStatus              string `csv:"fle-status"`
	FLEEncryptedFields     string `csv:"fle-encrypted-fields"`
	Port                   string `csv:"c-port"`
	TTFB                   string `csv:"time-to-first-byte"`
	EdgeDetailedResultType string `csv:"x-edge-detailed-result-type"`
	ContentType            string `csv:"sc-content-type"`
	ContentLen             string `csv:"sc-content-len"`
	RangeStart             string `csv:"sc-range-start"`
	RangeEnd               string `csv:"sc-range-end"`
}

// NewLine creates a new Line struct from a raw string.
func NewLine(record []string) (*Line, error) {
	return &Line{
		Date:                   record[Date],
		Time:                   record[Time],
		EdegeLocation:          record[EdegeLocation],
		SCBytes:                record[SCBytes],
		IP:                     record[IP],
		Method:                 record[Method],
		Host:                   record[Host],
		URI:                    record[URI],
		Status:                 record[Status],
		Referer:                record[Referer],
		UserAgent:              record[UserAgent],
		Query:                  record[Query],
		Cookie:                 record[Cookie],
		EdgeDefaultResultType:  record[EdgeDefaultResultType],
		EdgeRequestId:          record[EdgeRequestId],
		HostHeader:             record[HostHeader],
		Protocol:               record[Protocol],
		CSByes:                 record[CSByes],
		TimeTaken:              record[TimeTaken],
		ForwardedFor:           record[ForwardedFor],
		SSLProtocol:            record[SSLProtocol],
		SSLCypher:              record[SSLCypher],
		EdgeResponseResultType: record[EdgeResponseResultType],
		ProtocolVersion:        record[ProtocolVersion],
		FLEStatus:              record[FLEStatus],
		FLEEncryptedFields:     record[FLEEncryptedFields],
		Port:                   record[Port],
		TTFB:                   record[TTFB],
		EdgeDetailedResultType: record[EdgeDetailedResultType],
		ContentType:            record[ContentType],
		ContentLen:             record[ContentLen],
		RangeStart:             record[RangeStart],
		RangeEnd:               record[RangeEnd],
	}, nil
}
