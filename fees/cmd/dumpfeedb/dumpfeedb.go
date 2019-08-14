// Copyright (c) 2018 The Utopia developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Tool dumpfeedb can be used to dump the internal state of the buckets of an
// estimator's feedb so that it can be externally analyzed.
package main

import (
	"fmt"
	"os"
	"path"

	"github.com/UtopiaCoinOrg/ucd/ucutil"
	"github.com/UtopiaCoinOrg/ucd/fees"
	flags "github.com/jessevdk/go-flags"
)

type config struct {
	DB string `short:"b" long:"db" description:"Path to fee database"`
}

func main() {
	cfg := config{
		DB: path.Join(ucutil.AppDataDir("ucd", false), "data", "mainnet", "feesdb"),
	}

	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			parser.WriteHelp(os.Stderr)
		}
		return
	}

	ecfg := fees.EstimatorConfig{
		DatabaseFile:         cfg.DB,
		ReplaceBucketsOnLoad: true,
		MinBucketFee:         1,
		MaxBucketFee:         2,
		FeeRateStep:          fees.DefaultFeeRateStep,
	}
	est, err := fees.NewEstimator(&ecfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(est.DumpBuckets())
}
