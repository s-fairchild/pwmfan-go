# PWM Fan Driver for arm devices written in Go

## Audience

Intended to be used as a software PWM Fan controller that drives a GPIO PWM 3 wire fan (PWM, 5v, Ground).
Currently being used to drive my Raspberry Pi 4 [ICE Tower Fan](https://wiki.52pi.com/index.php?title=EP-0107) on [Arch Linux Arm](https://archlinuxarm.org/) OS.

![ICE Tower Fan](https://wiki.52pi.com/images/1/10/%E5%A1%94%E5%BC%8F%E9%A3%8E%E6%89%87-%E7%B2%BE-3.jpg)

## Installation Instructions

Currently this project is a work in progress.

Installation instructions/packaging hasn't been accomplished yet.

However to clone and run the code as is:

```bash
git clone https://github.com/s-fairchild/pwmfan-go.git
cd pwmfan-go
make build
```

This creates a binary formatted for `arm` architecture inside `pwmfan-go/build/` that can be ran from there.

## Configuration

The configuration file is `pwm-conf.json` and requires a PWM pin (preset to 18) and 4 temperature thresholds that can be modified.

Thresholds must be in descending order, and 4 must be provided.

The configuration file must be in the current working directory of the binary when executed. However this will change as I package the build to run as a service via systemd.
