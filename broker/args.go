package broker

import "flag"

var ActController = flag.String("c", "0", "a string")

var StorePath = flag.String("p", "../../def_store/", "a string")

var BrokerRestart = flag.String("r", "0", "a string")
