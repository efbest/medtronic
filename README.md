# medtronic

The `medtronic` package provides functions for communicating with
Medtronic insulin pumps using SPI-connected radio modules.

Decoding of messages to and from the pump is derived primarily from
[Ben West's pioneering "Decoding Carelink" work,](https://github.com/openaps/decocare)
along with [Pete Schwamb's code for RileyLink.](https://github.com/ps2/rileylink_ios)

### Radio configuration

The `medtronic` package and any programs that use it must be built
with the appropriate Go build tag for the radio. Currently supported radios:

* `-tags cc1101` for a [CC1101 radio module](http://www.ti.com/product/CC1101)
* `-tags cc111x` for a [CC1110 or CC1111 radio module](http://www.ti.com/product/cc1110-cc1111)
  flashed with [`subg_rfspy` firmware](https://github.com/ps2/subg_rfspy)
* `-tags rfm69` for a [RFM69HCW radio module](http://www.hoperf.com/rf_transceiver/modules/RFM69HCW.html)

### Utility programs

The `cmd` directory contains a number of command-line applications:

* `mdt` is a "Swiss army knife" application
(analogous to the the `openaps use pump ...` commands).
* `mmtune` scans for the best frequency with which to communicate with the pump.
* `pumphistory` retrieves pump history records and prints them.
* `sniff` listens for pump communications and prints the packets it receives.

### Documentation

<https://godoc.org/github.com/ecc1/medtronic>
