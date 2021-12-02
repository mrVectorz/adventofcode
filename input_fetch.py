#!/usr/bin/env python
# coding: utf-8
import os
import requests
from sys import argv
from sys import exit as sexit

class input_fetch():
    def __init__(self, args):
        self.args = args
        self.url = ""
        self.headers = ""
        self.base_url = "https://adventofcode.com/"
        self.arg_check()
        self.fetch(self.url)
    def arg_check(self):
        if len(self.args) != 4:
            print("Must give args: $session $year $day")
            sexit(2)
        else:
            self.year = self.args[2]
            self.day = self.args[3]
            self.headers = {'Cookie': 'session={}'.format(self.args[1])}
            self.url = "{}{}/day/{}/input".format(self.base_url,
                                        self.year,
                                        self.day
                                       )
            print(self.url)
    def fetch(self, url):
        r = requests.get(url, headers=self.headers)
        self.write_out(r.text)
    def write_out(self, day_input):
        with open("./{}/day_{}_input.txt".format(self.year, self.day), "w") as f:
            f.write(day_input)

if __name__ == "__main__":
  input_fetch(argv)
