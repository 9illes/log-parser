#!/bin/sh
zgrep -h "^[^#]" ../var/tmp/*.gz | sort > ../var/aggregate.log

