#!/bin/bash

set -e

egc cortexm/startup
egc cortexm/systick
egc delay
egc stm32/clock
egc stm32/flash
egc stm32/gpio
egc stm32/periph


