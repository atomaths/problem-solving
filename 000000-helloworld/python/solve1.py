# coding: utf-8
f = open('../1.in')
cnt = int(f.readline())
for x in range(cnt):
    print "Hello, %s!" % f.readline().strip()
