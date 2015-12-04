
import skilstak.colors as c
import time as t

def plain(message):
    "Normal message"
    print(c.clear + c.rc() + message + c.reset)
  
def nyan(message):
    "Nyan-esque message"
    print(c.clear)
    while True:
        print(c.rc() + message, end = " ")

def ayylmao():
    "ayy lmao"
    print(c.clear + c.rc() + "ayy..")
    t.sleep(1)
    while True:
        print(c.rc() + "LMAO", end = "")

def multi(message):
    "Multicoloredddddd"
    while True:
        print(c.clear + c.multi(message))
        t.sleep(.5)
