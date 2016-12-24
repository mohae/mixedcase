package mixedcase

import "testing"

var mixedTests = []struct {
	val       string
	MixedCase string
	mixedCase string
}{
	{"apple_oranges", "AppleOranges", "appleOranges"},
	{"catDog", "CatDog", "catDog"}, // if not snake, the first letter will still be in proper case
	{"DogCat", "DogCat", "dogCat"},
	{"Train_Engine", "TrainEngine", "trainEngine"},
	{"a_b_c", "ABC", "aBC"},
	{"a_cpu", "ACPU", "aCPU"},
	{"cpu_count", "CPUCount", "cpuCount"},
	{"Foo_bar_baz", "FooBarBaz", "fooBarBaz"},
	{"bAr_bAz_Biz", "BArBAzBiz", "bArBAzBiz"},
	{"9Bar", "NineBar", "nineBar"},
	{"ХЛЕБ", "ХЛЕБ", "хЛЕБ"},
	{"хлеб", "Хлеб", "хлеб"},
	{"хЛеБ", "ХЛеБ", "хЛеБ"},
	{"~Foo_Bar", "FooBar", "fooBar"},
	{"!Foo_Bar", "FooBar", "fooBar"},
	{"@Foo_Bar", "FooBar", "fooBar"},
	{"#Foo_Bar", "FooBar", "fooBar"},
	{"$Foo_Bar", "FooBar", "fooBar"},
	{"^Foo_Bar", "FooBar", "fooBar"},
	{"&Foo_Bar", "FooBar", "fooBar"},
	{"*!Foo_Bar", "FooBar", "fooBar"},
	{"-Foo_Bar", "FooBar", "fooBar"},
	{"_Foo_Bar", "FooBar", "fooBar"},
	{"=Foo_Bar", "FooBar", "fooBar"},
	{"+Foo_Bar", "FooBar", "fooBar"},
	{":Foo_Bar", "FooBar", "fooBar"},
	{".Foo_Bar", "FooBar", "fooBar"},
	{"<Foo_Bar", "FooBar", "fooBar"},
	{">Foo_Bar", "FooBar", "fooBar"},
}

func TestExported(t *testing.T) {
	for _, test := range mixedTests {
		v := Exported(test.val)
		if v != test.MixedCase {
			t.Errorf("%s: got %q; want %q", test.val, v, test.MixedCase)
		}
	}
}

func TestUnexported(t *testing.T) {
	for _, test := range mixedTests {
		v := Unexported(test.val)
		if v != test.mixedCase {
			t.Errorf("%s: got %q; want %q", test.val, v, test.mixedCase)
		}
	}
}

var initialismTests = []struct {
	val        string
	Initialism string
	initialism string
}{
	{"iDs", "iDs", "iDs"},    // an value that doesn't match should return original value
	{"xslt", "xslt", "xslt"}, // an value that doesn't match should return original value
	{"aPi", "API", "api"},
	{"ASCII", "ASCII", "ascii"},
	{"cpu", "CPU", "cpu"},
	{"Css", "CSS", "css"},
	{"DNS", "DNS", "dns"},
	{"EOF", "EOF", "eof"},
	{"gUID", "GUID", "guid"},
	{"Html", "HTML", "html"},
	{"http", "HTTP", "http"},
	{"HTTPS", "HTTPS", "https"},
	{"Id", "ID", "id"},
	{"ip", "IP", "ip"},
	{"Json", "JSON", "json"},
	{"lhs", "LHS", "lhs"},
	{"qps", "QPS", "qps"},
	{"RAM", "RAM", "ram"},
	{"rhs", "RHS", "rhs"},
	{"rpc", "RPC", "rpc"},
	{"SLA", "SLA", "sla"},
	{"SMTP", "SMTP", "smtp"},
	{"SNI", "SNI", "sni"},
	{"ssh", "SSH", "ssh"},
	{"TLS", "TLS", "tls"},
	{"ttl", "TTL", "ttl"},
	{"ui", "UI", "ui"},
	{"uid", "UID", "uid"},
	{"UUID", "UUID", "uuid"},
	{"URI", "URI", "uri"},
	{"url", "URL", "url"},
	{"UTF8", "UTF8", "utf8"},
	{"vm", "VM", "vm"},
	{"xml", "XML", "xml"},
}

func TestUpperInitialism(t *testing.T) {
	for _, test := range initialismTests {
		v := UpperInitialism(test.val)
		if v != test.Initialism {
			t.Errorf("%s: got %q; want %q", test.val, v, test.Initialism)
		}
	}
}

func TestLowerInitialism(t *testing.T) {
	for _, test := range initialismTests {
		v := LowerInitialism(test.val)
		if v != test.initialism {
			t.Errorf("%s: got %q; want %q", test.val, v, test.initialism)
		}
	}
}

var cleanTests = []struct {
	val string
	ndx int
}{
	{"alpha", 0},
	{"_alpha", 1},
	{"__alpha", 2},
	{"!alpha", 1},
	{"@alpha", 1},
	{"#alpha", 1},
	{"$alpha", 1},
	{"%alpha", 1},
	{"^alpha", 1},
	{"&alpha", 1},
	{"*alpha", 1},
	{"-alpha", 1},
	{"=alpha", 1},
	{"+alpha", 1},
	{":alpha", 1},
	{".alpha", 1},
	{"<alpha", 1},
	{">alpha", 1},
	{"!@#$_alpha", 5},
	{"__!@#_alpha", 6},
	{"@#$32alpha", 3},
}

func TestDiscardStart(t *testing.T) {
	for _, test := range cleanTests {
		ndx := discardStart(test.val)
		if ndx != test.ndx {
			t.Errorf("%s: got %d; want %d", test.val, ndx, test.ndx)
		}
	}
}

var numTests = []struct {
	val      string
	expected string
}{
	{"abc", "abc"},
	{"0abc", "ZeroAbc"},
	{"1abc", "OneAbc"},
	{"2abc", "TwoAbc"},
	{"3abc", "ThreeAbc"},
	{"4abc", "FourAbc"},
	{"5abc", "FiveAbc"},
	{"6abc", "SixAbc"},
	{"7abc", "SevenAbc"},
	{"8abc", "EightAbc"},
	{"9abc", "NineAbc"},
	{"_1abc", "_1abc"},
	{"ABC", "ABC"},
	{"0ABC", "ZeroABC"},
	{"1ABC", "OneABC"},
	{"2ABC", "TwoABC"},
	{"3ABC", "ThreeABC"},
	{"4ABC", "FourABC"},
	{"5ABC", "FiveABC"},
	{"6ABC", "SixABC"},
	{"7ABC", "SevenABC"},
	{"8ABC", "EightABC"},
	{"9ABC", "NineABC"},
	{"id", "id"},
	{"0id", "ZeroID"},
	{"1id", "OneID"},
	{"2id", "TwoID"},
	{"3id", "ThreeID"},
	{"4id", "FourID"},
	{"5id", "FiveID"},
	{"6id", "SixID"},
	{"7id", "SevenID"},
	{"8id", "EightID"},
	{"9id", "NineID"},
}

func TestNumToAlpha(t *testing.T) {
	for _, test := range numTests {
		v := numToAlpha(test.val)
		if v != test.expected {
			t.Errorf("%s: got %q; want %q", test.val, v, test.expected)
		}
	}
}

var runeTests = []struct {
	val      string
	expected string
}{
	{"", ""},
	{"hello", "hello"},
	{"Hello", "hello"},
	{"HELLO", "hELLO"},
	{"世界", "世界"},
	{"世", "世"},
	{"ХЛЕБ", "хЛЕБ"},
	{"хлеб", "хлеб"},
	{"Хлеб", "хлеб"},
	{"δέλτα", "δέλτα"},
	{"Δέλτα", "δέλτα"},
}

func TestLowerFirstRune(t *testing.T) {
	for _, test := range runeTests {
		v := LowerFirstRune(test.val)
		if v != test.expected {
			t.Errorf("%s: got %q; want %q", test.val, v, test.expected)
		}
	}
}
