package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
	"errors"
)

func EdlShotsToShots(edlName string, projectCode string, edls []*utility.EdlShot) ([]utility.Shot, error) {
	shots := []utility.Shot{}
	edlCode := *utility.GenerateCode(&edlName)
	for i := 0; i < len(edls); i++ {
		// TODO 调用C++程序,在原始素材中,根据始码抓取下一秒的一帧图片
		// 传入文件路径和抓取帧的时间点
		pic := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcUFhYaHSUfGhsjHBYWICwgIyYnKSopGR8tMC0oMCUoKSj/2wBDAQcHBwoIChMKChMoGhYaKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCj/wAARCAEeAf4DASIAAhEBAxEB/8QAHAAAAQUBAQEAAAAAAAAAAAAAAAECAwQFBgcI/8QASxAAAQMCAwUEBwUGBAMGBwAAAQACEQMEBSExBhJBUWETIjJxBxQzgZGh0SOSscHSFTRCUnLhFiRUgkRi8AgXQ1OT8TVFc4OissL/xAAYAQEBAQEBAAAAAAAAAAAAAAAAAQIDBP/EACERAQEBAQEBAAIDAQEBAAAAAAABEQIhMRJRAxNBMmEi/9oADAMBAAIRAxEAPwD5UQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhLCBEIQgEJYRCBEJYKIQIhLCAgRCWEiAQlhKgahLBRBQIhLBRBQIhCWECISwkQCE5JBQIhLCVA1CWEqBqEIQCEsIgoEQlhIgEIQgEISwgRCWEQgRCWEQgRCEvBAiEoCCgRCWEiAQlhEIEQlgohAiEsIhAiEsIhAiEsJEAhCWECIQhAISwiCg9wZsjs88EjC6P/AKlT9SqXOyeC0ySMNo7v9dT9S6OoHUagIjdS1QKtLKM1xnVVyw2YwSMsNo/fqfqUT9mcGYf/AIdSg/8AO/8AUtx7TTeQdCckjgHZKXqjEOzWCn/5dS++/wDUonbN4OHQcPpfef8AqW4DGR1UdUSJA0T8r+xj/wCG8H4YfS++/wDUmHZ7CAT/AJCl95/1Wu3RNqN+aflTGUdncHIysKX33/VM/wAP4RP7hS+8/wCq1gUxxIT8qYyauAYS1uVhSn+p/wBVD+w8LJ/caf3n/Va1dw3SqtN8k+cKflf2skV6OAYU52djT++/6q9S2awZzZOHUvvv/UrNBokStCm2G9F05tz1msC52ewhnhw+l99/6lnXuCYZTpkts6YPPef9V013+KycQ9kegS2rI521wuyfXh1swiY8Tvquks9nMIeG71hTOn8b/qsexH266/Dx3Wqy3SwyjsngRA3sNon/AH1M/wD8lHdbK4Iww3DKI/3v/UujtvB7lBeQR7lu/GZXL1NmsGDJGH0p/rf+pUP2Dhe9HqVP7zvquorD7Mzqsn/xCuP5VuSVg3ODYc2oA20YBP8AM76q9Y4BhVQDfsqZn/nf9U659t1lauGiI9y1LcOpiez2TwF/iwykf99T9S17fYrZt4zwigf/ALlX9SmsW6Lfs2d3Pku/E3652uYu9jNm2AlmEUB/vq/qXNX+zeDUqhDMNojOPHUP/wDS9GvRAIXJ4qO+ehXHu2XI1IyWbMYI5knDqUxPjf8AqRT2ZwUuzw6lmf53/qWzTEMA6JtPx+9SWiG02RwB5G9hlE/76n6lHtNhGyWA4U66uMJoOe47tKmKtWXu+9kOZ+q2W3NO1ovq1nhjGNLnOOUAcV5DtXjdTHsUdcOkWzAWUWHg36krpviSWsd7W1arnNpNpgkkMZMDpmZ/NNNKmNQCfMoLoGWiaX6QsbWg5jAcmj5ppa0DwifelLpGiQlNpgLWx4R80haI0SygAzwV2obuidAlLRyT0gHDgmhhaOSmtXUadT7Wi14njKYRBhIWyJTR2+D2uB39IgWNIVgJLd9+fzVz9hYUT+40gP6nfVcJZXVS2rNfTcWuGYjJd9g9+2+oB4I7RvjHIrN3/BM3Z3B92TYUve9/1UNTAcI4WFL7z/qtd5imNFVmSl6sWRnfsHCp/cqf3n/VAwDC/wDRUvvP+q1Q3jwTt3hBhc/yq4yhs/hX+hp/ed9U1+BYWP8Agaf3n/VbJbkSAq9SSrOrv0xmMwLCy7Oypfef9VI7Z/CoysaX33/VXqIAOanPBa2/tGK7AcMA/cacf1P+qb+xMM/0VP7z/qt11MOYdFSqN3TEqW0kZ5wTC/8ARUvvO+qP2HhkfudP4v8AqtAFKIlWW0Zb8Ew2P3On9531UTsGw7haM+876rZIVeqM9FqWozP2Ph/+kYP9zvqrdDAMNdJdZsgf8zvqrNJoLpWhTYG0x1V0Zf7Awr/RU/vO+qa/AcLH/BUvvP8AqtYjVRPIcVLaMz/D+GOEepsHk531Ults3hVar2brUCGl2T3Dl1Wh4WZqzhTd67d/QfxCm1XX3FMPYqDHGm8sdodFpUHhwg6FVb2iQS4aj8FzLdVbmnvMyiRmqbJzlXA7eGfDJVao3XyOKCOqND8Uw8uCmIDp5KuYmCYjRBG7I9Ehzan1B3ZCilFiNxiVHUf3hp1Cc85kZqCogZXMgqpRMVIzVp5yz4Kkwnt8j8kkG1bcFoM0noqNpm0ErQaO78l0jNU7nVZOJeycte61WPieVN0clL41GZhrftz5rrrEd1q5TCx9sAOf4rrbEd0TwVntK2aEbg5QoLz5QrFHw/JVrsEmOa3b45qFxPZu5rKbnUK17gDsTzWOPGSuN8jpFK5I7Ra+GjRY1yftls4YchlwC1yvV10uHt0nSYXQWvhWFh/Bb9vG6MtV6eL442KV/wAeULk8UJ3zPGF1t+MndQuRxXOp71w/l/6105+Jafsh1Cg3wDAOhUtN32TRHBc/tViJw7DXupu3a1Tus6SNVmbpYwdtMeddV3WFq/7Bh+0cDk4jh5fmuSec8uCQkuJc455/FNbLjOa1iTwh/NO6AJCM04ak+5F0gGs8EnGAnObrmnU2S4ZIFczdDeuqQM7sjRS1WkkQn1GgNaM8oHxT4iIsjI8YKYWiJ/6hXiwOzyyyVc04J+aauISN5m8BKVveE58k6lqBwT2U4eWxrkEMQ7kOkrVwS7da3THCQCY14Hgs94mJ1BzT6ctdI1lZtMeo0nNrUmPHhcAZlRmlGg0WXsteCrbuoOJlsOHkf7rac6Cp0SemtENAHBAzKWZSCASubRxChewclKXwAM0xzpPFWQQsHfy0BKlfoCeqGNMyUVOAW8rKQHLOIVaq2SVMPDqmbpclhqke6YIKkyhFdm7nkmsIKkuXFv8A4fnooa2QlSqK50AC1KhbRpLh8VokcNAFWsGcSrbhxPvWtRDUIGnFQsE/kn1T3wApKTQACY0Wdm4qJ4MZq3hGV2f/AKZ/EKrWcM4OZKmwkk3bif5D+IRHUW7oAAVt8VKeeo+ay7d8xmr9F+YkiFieqzq7ezqnkSo6jd4dRmrt4zeZlqM/cqId5II2ugQVXrDdMjRSu7rvNR1ILfIoIplpB5qKpoYTiYhRPdmc0WI6hULyIJ4wpHuAz4qGtmMuKLERfvNIyzyTKVIuqznKQ0najiprV0OgqxGpbNLWCdVfb4B5Krb5gK6MmrcZqhX8UEcVjYrlTdwlbNxqsTFj3HT1Wd9a5U8K9uPcutsVyeFZ1czxXX2TcmrUK1qIhqr3euqt0xDAql7l8Qt9fGFG5P2RHvWMPGYWtd+zWQw98+S4ukUrj2ui2sM0HNY1f2vvWzhmoCvJ06ewyAK37Y9wLBsAIC3rfwe5ev8Ajcb6q30ZieC5DFfajoV1t6dVyOKH7WOq8/8ALf8A6dOSt9iJ4BeY7Z35u8VNMEllLujz+q9Fv7kWmG1qztKbC4e7+68cr1HVq1So8y5ziT71OfYEc7lonsHdkeXuUUSeqsNyaPJazIhmhKVo0nUFI46nLglE7xyQO1cRGSlpjdknyCgpqyGue0boJEHQcUCNO8SQNE+pmRlpCSmxzWSR8k1mbwOJKl9WLbMmHzk+9M3QT8VK5pY2CDzUW8A8A8RCCqW7lQ/JWH05LKgykiVHcA74IOoVmh37Yt1IzUxdNr05qODW5ObI4ZqBrZzB0K0CJ7NxPQzwyVcsDS8TEZFXBewm4daXTKgOUwRzBC7US4Aggg/NcDRgMbJOWS7HBrntbKnvHw5Z56LNiL7QYEpzoiJRlwTHFSzQhBJ1VilTykpKLQQCVM4wMikKYWxwUNRu8RE5KwG705FDmwASOi19+nxWaIGacCAc/wAEOImEhHJL4G3LQRkqLcnbvVXajjukEKjUydKx19WYkHFNcN4jySs0yTo78LWeIuWgDaadUIAPVFEbrAEyqcir8gq1Cd9x6p5qHdhRPOZTmCYWIYdEiSrWGgCucv4D+IVZ0NEJ1vU+1PktXwbGG1g9jXStVjhAg6rj9nbzfpgE8V1FN5yI4LC1aqmWH/qVmEbryDpJKvB8jNU7oZghEQ1tARGSql2vMqdxJaQqjjmUDKpULnZ5FOqHWFTc8z1QOqkkwEjSZzClpQ7MgZmFLWphrQeaLpg3TTzAVemR2pA1Cc9wAMHgqlvJr5SrDW/baBXx4ZPJUrYQ1sK5/Dmtys1SuPEVg4t4XLeuPF5LAxjwlZaiHBvaBddY8AuSwUfae9dfZDQrXM9K2KXhVK91Hmr1LRUr0CR5rp1cjnPrOu/Z/NZDPER0K1732fmsdnicei4WOsU6+dULawvUdM1i1farawvgnP06/bqcP8LYC26Ebo8li4f4Rp8FtUR3V7OPHGqF+cjC5PEJNeF1eIaHouTvz9vxlef+X66csTbGt2eC1GNneqxTbHUifwXmDxBPn+C9M2mpG4ZaUw4A75cfcCV5xdx29WBlvH3KcVaib4hHNTCczmiyo+sVhTEAwTJ6LescItXkdtcvPMMbK1Urn4JOSt2lhcXTooUX1J/laTHOV3GF7O4FVewXF1Wb/WA0EdSM16Ng2A2VlatGH0qfZcHNIM/3UtS3Hg9PCbsXrbR1Cp2782tDdZ/JepbJbCU7elSuLt7nvrU+9TLYDSeY1JXcm1oksc+hSL2HuktEg8xxlLVrdiAGAF5yaOXuUt1i9OSuthcKou3qjj2bnGWg97OdOahHo2s6v2ljTrB0GXVXboHUA5+a6HEMUtMDpuusRqtNd4zc86dGj8lxmI+k+4q1nNwu1L2zAqVyQJ57o+oRZrWq+jWpXpllS5tmOa0tD2AmdCJ+a5i/9GGP0i91syhdU6YLg6nVA3vccyctFfscf2rxKpNPEHscdKdrQbA85C6nDH7e0H7xt6d5bkgltzSphxEcC0h0/FXKSvD72lUoVDSr0n06jSQ5jwQQR5p1k6HRzEL2TbGytMZwfEK1zhz7PGLaiaj6VVsPaAfG138TOExkcjGU+KWzjTuHNMyDCRqVoDMECMzKa4b1RukuH4f+yV0bwcOGqU5hruDXT/7K6ptNpLYGoI/BdJs+6WuYTke9ksCkM3NBGYjyWvgb9yqwTlp8FLIT9OiB3dVNSp72ZUTxlIT7arDgD5LBqxAaITWy5x5IJDjIOqkDd1vuzWp6WlBDWjMZqOq8EQFDWqbuQlRBxcDKW5EhHO7xg6oDjxkpjxDuPNO0EkpLq3w8DeaZ5KtWZBz0U1N2afVbvMMa+SnUSeKTDpKlbEqM5GZzSscJCkrVaDIgKKqcimioY46pHOkHNaqIYlxU9Noa2VCzN3vUzvDA5KSf6GVj3SqzX7hnpCkqzunNVL87tuNPGPwKWkjNwK5NOqGnQFd9ZVQ+k0iZAXmFnULLgHIZwu+wSvv0WgxoucrfUbW91UVfvMI5JZSPjdM8lZNYU3OAGqpVXZkSpLh+60xwWe+pvGVcEjnnmUlJm8YGsJWs3hkc1PSYGHPlGqsmBvZBuc5aqO6q90wm3dYgkDVQMdvyHIK2+7ezJVy1a0vBGozTKlEa9FJZNIJyTPRs0NArR0VWiNArXBakKo3HiK5/GTDTzW/XPfPmuexk5HNZ/wBWHYIBvAx5rrbICAAuUwQCQfcutsx4YWub6z01aZ7o8lQu/aHNXWzuA9FRu3d8rfXxnn6o3p+zhYzP4lr3/sznwlZNOIdmuVdZFKp7UDqt3CtAsJ/tR7lu4XECCnKdfHU2AyHktqj4I6LGsuHkFr0T3fPRevlxqhiEd4rk7794PJdXf8Vyl6ftzC838v8A068sLHXBnbVyDFGiY83Dh8l5lUdLjOsr0HaysBhNRuW9VrRy7rR/7Beeu16ynK2LFhvmtFOAS0gnkOauvw68J7Sm9ziM8iRCpWFQU60u4iJXYYU4boM/Nau6ljBtcaxXDX7orvc3jTrt32n3GV3exe2NCpctoVSzD7h5gEEm2qH/AJpzYeuY5wqVzY291TIe1pHmFy2J4KbcudRf3ZORHDTUcFm5fKmPopry5nfZuPGTm8jyWVjmIU8Hw26xCqw1OwplwZOpkQJ8yFmbG47aYhhFpb0qr3V7aiyi/tBDnFoA3vLL4LqH4RTxG0qMuabatu8btRjhMj5e74hZ2T6xZlfOt3dX+0mLuq3JdVrvMtaJ3WNngOXVeg7JbIscGGu0OGRM8foPmrFphVnQxqvSw+0Za2bXwwSSXxxLnEk8ctPxWtiW2GE4C3sA8164yLKf10Wfy25G7L/jt8Cw+1w8AU6TARkMgumpXDQ3OF4JW9LFQO/y9gA0fz1AFNg3pHx3E7ns7TB33omIt3Hu+Z3YXWXxi82vXNoTb1LOu+oxpIp1ADy3mkGPlPzXydctFK8IbplHWCvcNo8Ux+5wGsG4f6uXNO+Xva8s8oyJ68F4nesLKzC7XSfes/ltb55yepjmBPDJObm3qo3GGmOOalECqB/C4SPeFpUtPOtPEt16q9YEsLXR4XSY5aKiyW1mEaFXbaN5/wA1f/COsY4PZ7gmuG6ZHEqKzfvU2E8lYqiQPP8AJZsE1vnuypazw1p15JlrwS1hvA+aT4ll3xUcS90lOa2NEAd6FKBksW+rJ4iNOdSoiDMZ5K0onU85CktXELcjKlB7ufJNqANA80jXS3zXSXxFO5duugJLc75HRMvZnMZItHwYUwXiICa6c4CdvSEgEg5ohlM5idVOTkoWtO9kNFKBlorBHUbIIVDFWxaN/rH4Fajmqhiw/wAqI/nH4Fc761HINPeESuz2br90CdPmuPpU3VJ3ADHVdBge/Tc3eBACzrpcx3IIiZVWvWABHIJBWmlI0hUq7ySRw0K1HPPVW5qkuyOSgptLtNAisJfAjJPp90SOS1DFqlDWyeSp17hwqABBr5kZZJpp743gMyrpiVjRVgwJSGluSYzT7dpbA6qxWIiTropTFcnuZpbSN/8AumVXZQNFJZA7/TiiNahqFOeCgoZR7lYOi1PSqNfxFc5jJHDmuir+N3mucxg96OsLP+rFnAtF1ljqNFyuBaDkurs4kK8xnq+NL+EeSzLvxe9aQPcELPuo7Q+a6dfGZ5Wdf+yKxw/dDlsX/svcufqEgOXKx25/ZC7eqgg8Qt/DAA0c1zdGTVHVdNhoyaByTmZU6dFZk5LXpHuarItOC1KRloXq5rhfqnfu1zXL3RBuPMrexaoWgwVzFeoXVHE6hef+T668uM2vqB3Z0wRA3jHm4n8lyBB4nNdHtW4G7MH+KPcB/dc6/U+acrYYtKnc3XYNFKWjIbwyWeei6XC2MdRaHNBEBapJWWKeJ1TLe3cfNPfSxSi2arnsGnedK6y2tKAAIaQddSoMTYwMJDRI0zU3/wALM+qOydzXo4pbVmkscXgOjiDHBfTmzlMV8MdvT3gvmXBWk4jbgD+MfivqTZZobhFM5TACz+Mc+/GBiez9N9GtTFPde5pDXtyIkarzxvo5tnOca9APcSSS57vqvarphcTCyK9ENcSRmp+H6WdXHnlhsDh1BwccOtZBkGpL59xyXY2FjSoUhTO4KY0psbutHk0QFNXduNyAVMVyKmenBanP7rrxxeol2jpN/ZFUNAADTA9y+cMcaGVxPAkfivoPaK9aMLqNBzLTK+fsf71wSOJJCvU/TN5vNyqrASyCntJ7NmckZJKBmmJTmgAEQFJUWAYLSeBy8lZouioSDwKq70AdFLTyqN5Lcg6bD3F1MZcifgr9QHcEaLNwaCyDwceHDJbRaNwTCzYG2xMDmpKo+aLcZlSVBln5KfIKoHxTwEpSrlVlNjzQQnJChqtXbkQEUKYLQTCfV8KW3B+JW+EtUb+jrA6qnbiHyStW+GUHqsul4wBqtjQjLolboU5jZZMJWgx5KYhtMAH4KR+gI0UbQO0IU0d2OKq1GdFn4x+6N08Y/ArQes/GCPVW/wBY/ArlfFizhGFUhQAc1pPEwrRs2U3Ddj3cVi22OGkzdbEfFTtxdrjJiSea8/U7tdJ1rXeSAAOXBNdDgZVSnd9s2QR8VHWuCwkAmYXfi+esX2+C4hriZVRtYl0EmE2tW7TUmUtKlMHPVbl1n4simS0EKxQbAAKWm0BsdEpeG6IupJAnIKCq/wCSaXSdU10nihprpOqsWXiBzVcSdFcsxmhWnRU5OQUVEaGQpSOBWozWfXPeceq5rGZ310lzq4yuZxYzWhZz1Y0MCBgZLqrPULmcD8IXUWeoWufqdL/8OSzbkneK0yIbPRZdyZetd3Ik9UsQ9n7lhuYSHHLitu/9n7lkkdxw6LlrpzVKjHbZ8Culw/wtjkubo+3ldJh/hAW+funVdBZnJarPCI1WTbaBaTcgOi9HN8cazsUbLSQAuTvwGuc45AZldbf+F0grkMYO6KkakRz1Xm7/AOnXjM9ee7QVC+8IjiT8Ssd2p81pYu7evXkaAkDyBWY/xFXkpp1W/hdY7jQDJj5rnzqVfwusabyDpqtVZcrr6FU7gg8FVxGoC2CTmoqV1Ra2e0aMuap3VwLh/cktHFZta76l9jT2aZ2mMWzRPjBX09s1lhLAeS+Z9kXRjVDoJ9y+l9l3U/UmB7wBH5K64d1Zr1xTcAQT81nX1RoO85pAOhiR8Qn4pi+GWVWLm6DXOMBoaXZ8sgp7CrTuZfSO9ScI5SPeudt3xiWxzt1VYWyDkeKzalZrCSIngujxPAbGq41GtqUXnMmk8t+QyXmPpJrN2dsWOoYhcuuazjuUn7p7gMOcYExMN6nyKu35Xfj+SyYbtdizWWr6Ydm7KJXk+IPNWsSNASrt9cXr7KlcXtUvfV7wbGgIy9/FZ7jvAEgq7dwt2i3JiORU8cVBRHIRKs5ZTxWpELz5BSscSZ1OuiiLZBjkn0uR5rco6TA3HdM8xw5hb7o7ISucwV2oPMH5f2XQz9mAeGax1fSnWx78KWrw8lXoO7+Smq8SpfYiLil6JpcAgFcqp2XBIUkoLkXEVU5J1roo6swY5It3GVvm4zYW9iPfCy2N/wAwMuK1LsyOqoD2oPCVd9VqU2/ZxmmREwpaJJp8dOKjdxWtRDPfUjjkoSYefMJ5I3QgY8mFSxKia9AM5PB+R+qt1DAMKtWrbgnqAuXV/TUcRvOEGSntrP4HNNdE5IaO8FpHRYQ5zmQZU1zIqFWMCoAURlmROimvKIL8oyUk9WVnNZLhE9Vo0ButzjJQBoacxnCmaTAiPirmFWt+BoFESSiTASAEnVEA1ySkBAHBKdCi4jEQFetOEKiOCu2c5FCtOllEKZxyChpKR/hK1GWfdR3lzGJn7ZdNc6HouXxP23vWd2tRtYJkwRrkumsdRPJc1go7g8guns9ZWuU6XneBZdwO8Vqu8BWTcnvkrXfxjlQvo3Nc1mVCNxxjXktK9mDkst/gd0C42+OsU6R+3XS4edOS5mifthK6XDs4WuU6dFaAd1aTQNxZtpwhaTT3IXq5cqz7/R0ri8cdute7gCF2t/o5cBtLULaT45/kfovN3/06c/HAXmdZ0qm8CYCuXR3qhPVU6niKQqJ0SU+3qdlWa+JAOfVMcE3kukF+u0tqdpSMgd9sZ5cleoVW1WAtjPUclRsXGow0st9p3mE8+SZUou3i6jI5jiP7LOf4Np2IVbIU6rHw+n7PLn+IW3b+kTaP1R9O2fSpsptl9SnSktBIEkmQOS4Ytqb4FQOnqfqvUthLvZ7Cdn6lDEby2fc32dzT3S+GCd1ndGesnqeiXwzWhsDt1Vt67hjNSteWNV286O++i/TeE6iNRkOIzkL1hm2ezDaUjHbADXdL3Bw82RM+6V8+3Nha2eIufs7Vub60qTDBb1G1KU9S2HD/AK6rTpYXj18wU6VjXaDADqzgyPzTOft+s3i3/Ho21HpRwazpPGGtrYhXiG9x1KkOpc7vEeQn8V5HaOvtscfrXuKPL6LHB9YxDYHhptHAZ6cBJ1zPY4P6OW7xuNoqzqnEUaNQsaP6neI+QjzVjaCrZ4bh5tcOoU7a1pgwGDdH1nmcypbnxZzI862oresXwpN0BJyyWW8AaaBWy11SpUr1Bm4kjynRVq4iT71iT9qbQHJSujXkmUdBzSuPdPn8V0hUocBHOc06nG8JiIUTyZMdCnU8zPEBaNb2CkB8GJj810W9FPTULmMFd9s3WIJ/NdLTzpjqs9QFFxDtFM4mFExueXNSOWZ8EZlAmOKHHPRDSsVcKSmBxKc4hMJHvUUO8PRFCN7PWUx7u6QComVC0kzwWufrNT3ZAAVAn7RS16u+DJ1UABkElbxGtbOmnBTajwMuKZQeBTjiq9xUkkJgH1JfI+RT98QqTo3py+KHVEwT1H5ZLNxQn1dv9Q/NSmpmRwVfE3j1Zon+IfgVixqVzLiU9mbmgprinUvaCea1LU+O/wAFogWzXHXdUF08NrvAVvCHhtm2Rw/JZdzU3rh8zElKk9Nec1K3MKvq5WacQFKqQCU4DRBGSe0ZpVN5IdpqlJTSTBUNIr1iAVn8QtCzyCmlaNLMhSPB3VHQ1HkpKngK6Rhm3GbSuXxLOuB1XTXToBjVcxfn/MgdVhuOgwb2Y5QF0lmdMlzuDiKYPQLorNb5Z6W3nuz0WbW9oVpP8B8ll1z9oSt9eRIpX4G5PvWRVPcK1sRd3SPcsit7MxouPTpzFSh7YLpsO0HkuYt/bBdPhujZV5SuhtNB5LQZoFn22g8ldacpXp5cqq3vhdOglebbWVRvPaCIn4AT/deh4jU3aTjlMFeZbVOBruA1MCF5+rvTpHL14nLVUn9YV2vEqmZnJII3CCojqrDgCMzkoXBalIGPLHhzDBGi0nO7akLimcxAqDqsvqrFnX7KpBzY7JwVs2DZsaoD2kta4dRK7rAcRtGtaHFu8AMg0n5Lz2gBTrBuZbEjyWxfGkMGe5xLagzY4GCD9CMlzzfGp1+L1m1xO3Y3fcHNaM5fDB7y7ILoratcbgNK3awcyZ/BeJeiahZV9s8POIlpMu7IVMwa0Hck+enWF9CGmG0yAIPFT8Mvidfy34wb4VHN3q7y4jMAZCfJefbTA3lc27PZsM1D+S7/AB2oWU9xkdo7IDquDxyLSg9oIL3HM8yl2eOcv7cZfgdt2bAIBzWVePG+WN1Cnxa/FtvU6ZBqnMnWP7rJtyXNJPHM+asl+1pdbO6DzSnRIydxqCth5zAnVOoalMHBPoiZjgVqLGvg/t2ieBXU0c6YIXDUaz6WdMlpHIqelj19Q7u+1zRoHNB+almldmzxHknuiJXJUtpqzfHSY7LmRKtjaek4d6g4GODvqpebmGtt7u8m78ZysQY7bOJJDx7lZt76jXH2bwekwud5sXV2pViYKhNbk75plYjcMFQNMulJz+zdWxUlve1RvAhQQ4gwEtJrs5WpMS+mOneOR1RLsjJVrdEdVG5vwWjDmPIAJUNVxOYhMrPgQFXL+qEiQkjRJORlJvpHVJ10QxEXxooL1/2Qz/iH4KZ+k5Kpd+yE8x+alMY7tU+n7RvmEw6pWO77ecrI7rCnEWgHRUqg+0OatYWSLYZ5Qqzs6zvNCGtB3oVqmFXa3v8AFWQputYfrCeOaaOCcIiOKgQ6pj9J55J5THxuoYaBkFoWeQzWeMoWjaZjPRErRoaJ1XwFJRCWrG6ujMZV3MHTmuZvf3kHqulvD3SuZujNyB1WK3HS4T7Ie5dDZ8Fz2Ez2Tfcuhsxot8Xax0tP9mVlVfGVrPjczPBZVU98jjK13ciRQxDQwsmt7MytW/0Kyq3slxt105Vbf2w6LpsNGi5q1H22fkunw3QRwWuZ7qdN62bkFdEwAVUt9ByhXBmBmvTy5VlYrPYunQLy7aaoH4g4DgM+q9Mx2s2nbPLicoyHEyDl1XkuK1S+vUeYlziBn1K8/U/+nSMqu7WDxUB+alqZlRZjJRcI7RRJ7yITCtQhruianHkmrRi9ZViXNY6JacitLFHPdhwIndGvlIWACQQRqujsCy+w6tRcR2gbIE8Uz3UsQYW8i3rVg4tqUwXNcDoWwQfNfVFOua9BtWo0te9jHOHIloJHxJXyxgNI17qlbNBBq1qdIjq5wavpi9vqFrRrvqVWMotcSXkgBoEjyjgpbIx1rMxh7KYqVHxI+S8c2yx5lO4qU6RDq8xHBnU9eiu7e7etvHutMFcTTkh9xz4d36/DmvNXO3iSZJJ4lSc77SSz6Wo81HlziSSTqVctR9mOUKi0E6LQoiGQBnED3q1uLY8I6JCc45pXHdGRyAUG8TVZA1BVMTuMRJ8k6g/vPB1+igru7g8vjB1TKTyK0ka/RNFt790Zyq1VxLgnvMtBjMAGFCcxmpvoTePApd88YTTkctEZK6YXtDKfTuHMeCDGfAwoUhTUx1uDXbbsdnUP2gE+a1hRYB5LhsPrut6zHsJlpB9y7e3rCvRbUb/EJUsE26I0CaYBJCRzoGZUTqgH9lMVIXE6fimudxJyUDqw8lA+4aBrn1VElUjPqoCRooK1yNBw1VX1oTqVm2C+SI1TXHkVVFyDnklFYHOUlEznHd1+SrXLu5pxTnVAQYOagruJYPNF1nnVDR3x5oOqVnjb5hSbqOzw0/5YCeCib43Rrmn4dItwo2fxJ+yHM8SnGoUFPX3qwFJWpTpzGSemgJeCgTimVNE8qKp4QgVuq07MZA9Fmt1C1LUd1EvxfpDLqkreEpWJK3hPkujLIvNCubuf3oea6S88JXNVs7wcpWK26jCR9m0FdBawIgarBwv2Tea6Cz08l05YtWKg7pWTV8ZPVa1UDcMFZbx3z5p2k+s3ENCsyv7I+a1MQIzWXW9muVdeVa09qumw3guZtMqo8102H/wha5+p037ciAOOSuOcGU5JGio0iQ2fesXaXaK0wqm9leo51yB3aDRn/uJ0H4r0y5PXHNUNsL5vZljXBocMj/K0DN35BeaXdcVqznAANHh6BWsXxeviVUvd3QdRzP05cAFmHJsFefr266whMmeSYTnJQTKa4qSKaeKQ6JyjK3IlIdU0pfJEKrKRPpVH0nh9Nxa4cQUzNCqa0LS+fbvdWpvcy4a5r2PaB3XAzP8A1xhOxHGsRxM/5+9uLgcnvJHw0WahTEzTvwSBueqRKJ4KmVNS6q5T4eYVOkCNfNXaXibyhZD6rsiOJyCYSRcNjhCjqvJfE5DNNLvtARzhKJ6zhAyEthVmuMieCe8ySVEIk55qW0WTUMgjmnaj3Ks2ZzmFZpnKCMuaoRwEZ6qM8FM7IcwojqgRCUJY81TRTJB6LocGuuyO44nccAPeFzwyWjZv7vUKyazXQVbloMyc1TqXccRCoVqjhkNFVe9xXLrvLjpzzb6vVbzk7PzVZ9w58nn1VYyfJLwyWPytbnMiRznO1KiM8CU7eyQ4qf61JhoqOGkI7ZwOmSagceqbYl5lSesZZ6pRX3goCMkkZqys3mJjCWk3erMAmZSHXMBTWLZu6Y6rrJ+3N19qw07YdBKr0o706K+1u7ZunWFQpQGk8VCHU4n5qdqgpanJTtWbDakThmmpWqLLSO1UT9AOqkOqjfmint8QWpajuysumJdlzWtajuBIVcp/gm1j3SnUxkmXHg963/jDJvvAT0XNOl13J5wukvvZkLnBnd5c1PG3VYYPs2roLIT8FgYb4BquhssgtcMdJ6o7p00WVUJDiFq1vAsqr41e05ZuInIysqu7uFaWI8Qsu4J3IAXKzHXlFZ51V0NO6oWdv291VZSpN/icdTyA1nouSrX9PDqTqtQBztGtmN4/TmuVxLE7q/rdpcVSd2d1ugb5Dh56rfMTr12WP7eVntdQwZhos0NxUEvP9I0A65lcRUuHVarqlVxe90kkmSSevNVS4zOqSc+K39ZWTUH/AEUw1JlQk+aAVMErn5JsyQmSlByVyBXE6JpKHGUqBqEIQ0IQlyQIiEuXBKgaBmpGxOSaGlSNaQJhA4cOUq3k1k9MlTLt3Q5p/a97dOsAFA4iXSfJI6d4+eSQEyMk/KJCGwlNu86DzTjT3Sc9U6nk4dE55l3RQRBpAzCkaYGZSzLOGSYmBZJQkHzS+aIEo0SFOGioQDMBXLU556wqo0mVPRMOBjkrEWSS5pk5t4KI65aSlcS10j3pHdFy/lmXY6/xdbMMPWErdE0zOac1cq6EeIOabCe+VHPey1Qp8ZKM5FSzkonKhEQEQjiiVIdVZw798p/FVoyVrDf3ti7uLsnk+qmTwVCn7OfL3q3VP+VIOkKqz2axCHU+KnbooqWpUzdApaHBK1KBlMoAyKzqmHXqo6niaOalOpUL/G3zVEtLxZ81rWvgCyaXi9617Yd0Kwq3T8IUdz4PepGeFR3PgWr8ZY9+YYfL4rnKed3710F+e45c9Rzu/es/I1/jrsOHcELftJj3LDw0fZjmt+18IXTj5rPV/wAPq5N9yy3wahWtXHc9yynDvk8k7OWTiXFZVy7dpEkgCCT0C1sSXJbUXfYWXZNJD6vdnoNfyXPLa3Lkc1id267u3vkhgyaOQVJ3JKdUhGq6/GdIhCEAhCEDkkIbqn80DIRCfCRA3NInpInMIG5pYTkAIGQeSexhOvmnNanwECNACWo+AIiQmvdlCjJk5IHUgXvLjEAJlR0vmdVPu7tGBqYJTHMyQLb1P4XTy8lYIjjI1VDNp8lfoPFRpB1hKlgBzlEmUpbzSIFnLLRCSInqlGZz1QCUJEs8D7kDkD5JBwyySgZIFOXkU+kYIBy80xxGQSMMPg6K/D6uvBk+SRhlmaUx3fL5pjMnQnc/KJzcoOplK2Ej9TCRq8t/T0z098ScwoiYKe45Jh1KKfkmPOZS5wmO1zQwBA6pAlCsjNSmOKtYZ+9CFh+v1f5WfAqW3xWvQqb7GUieoMfiu/4uFrv6zv8ALjPJRMEMXJu2ovHM3TRto/pd+pN/xPeAR2Vt9131U/GkrsqOhU4XEt2pvW6Ubb7rvqlG1l9/5Fr9136lLxV13AGXuThELh/8XX//AJFr9x31R/i++07C0+479Sf1012btVE4faBcgdrL4/8Ag2v3XfVNO1V6XA9jaz/S79Sn9dXXa0vEVs23hHLJeZN2rvWkkUbX7rvqrNPbnEmaW9l9x36lZxYlr09gyiFFc+zMrzkbf4mP+GsfuP8A1JtTbzEniDb2Mf0O/UtXnZiOsxLwFc/Qj1weayq+119WEOo2oB5Nd9VSZjtyypvinQmeR+qz+Fxdep4d7Ie5b1oe4vIKO2uI0QAy3syOrHfqVyn6RcVpiBa2EdWP/UunMz6l9esVvBIWY7IleeP9I+LOEG1w/wD9N/6lXO3+Jk521j9x/wCpTqakjs8VMAgDmvNtobn1nEHBubKXcH5q9cbY4hXJ3qNo3q1jvzcucdUc5xJOZJJKk5sa3wplCj3yjeK1gkSAJm8eiXePIJgcQEGEzePRLJ5JgcE8cVDvdEu+7p8EwTt0zzSHpCjFQjQBN3zyCmUTREIOWSi3z0R2jp4JlEzQM0sKEVCOA+CO1dyHwTKJxAPRMc6ZhROqOOsJN4zKuBSSVNbU9+oCdBqq8lS07h1NpDWtz5pguu1PIKMsHxVYXL/+X4I9ZfyaplQ+rT4wMkyk8035eRQbhx4NUbnkkkxmqrSHeExlqkIVJly9oAEQOYT/AFt3JnzUyoslKdFUNy/k35o9Zdyb80yi0OsJepCqesvjwt+BR6y/k35plF1sEp+kznCoC7qDQM+BTvXakRDPgUyi4+I6lMbk7hoqhu6h1DPgU0XL5mGqjYb4WOlI+Q+Qs0X1UM3QGR5FK6/qu1az4FWVnGk4iJgSU2VnevVI8LPgUnrtT+VnwXLrjb46895PWk45JpiZWf69U/kZ8Enr1T+VnwU/rrX9jSnJRuOqo+uv5NSetv8A5Wp/XV/OLpJAlOaZVA3L+TfmlbdvH8LSp/XU/OVUQhC7uQQhCAQhCATk1CAQhCAQhCAQhCAQhCAQhCASgSkV7Dbl9lf212ynRquoVG1RTrMD6bi0zDmnIjLMHUIPbWbIbF03+j/ZbHcMxpu0eJ21GpcV7CsynBuKp3O1a9rp3GxpEDI9OO9Juz2FYLszsnUwtgc+u7EWVbk+K4FK7exjnCYndAGS9LwHbDaGw2Xvts/SNa2tapTqPr7Om8YaV0bp4iKIEH1dodvEHu6RmsG52iwmz9E+w2G7VYb+0cJv7XEKjn0YbdW1f1l27VpPPGYlp7pHOIIZ2wGwuyeP7LWt9fuvGXpL2VgMesrVu8Dwp1GFwERqea7vbb0U7CWmP3u5SuLSlTZTd2NDaGyoho7JpMU6rS8EzOZzmRAICdsTtZiOC+jvE7vCtqa+1GIYNZW1GhhlCzDLa3NU9mwF25v1izMxkMs+m9hW2W0F7gjcWvq20NXHsGtadPFdlaFrTo161QZtuXOc3tG0nNLd4MbIJygZoPLvQfsVs/tNhe0d5e4Y/EXWt1QpW9Ko6u8spPFQkltuN4u7oE+HXRem4V6JtlLrE7ahc7IsZRqVGte4UsVZug8d543R74C4/wBEtXHcVwjG726wnZ8NrvuBh9TErKj21/iVRxe2j2jx34706ZgNy4dhi91a2NhgP7XpYVgWJ3OG0ri6s/8ABhvC2qXODiXNI3TlG5EiJ4oPA9ja+zeEbe3NptBgDsdwl9d9s2myq/tKQ3yA9gYRvugDLjnovaqewGFUmuxC82BwV2zjKYrm8tql++5qMJgNba73aNeYjvQ0GTvEDPz3BNrcRw3azaKk2vtFb4RXq9oBszZNsXPeAGseabmk02lgJ3Z1966muzZ+nsPY40219Ibb9+K12G9p1GDETFGnk+puyaUaDgd5B4Ztk+2q7SX9Wxwh+DWjqpNGxeXE0W8AS7OefVdzsRsla4t6FNvMZfZU62I2NW2NrWIO/TaHTVgzEbjpPl5KLamltH6Q9rsPwrCW7U4m1rP8tTxxwfVpb0B7i4ANayQ2SeWa6S52wwj0e4ts9srh9Sni+DYb24x+pQM07+rcMFOs1meYYyGg5SRmg8JdqUi730m7BXGxt/RuLaqMQ2dvx2uG4nTE061M5gEjIPA1HvCwdotm7/AKOFVb/wBXNPErRt5buo1W1AabiW5kHIgtII4IMJoz5r2LA/RVZ2e1uA4djm1GBC+uqtrVdhhZcPe9lUtcKcinuyWmNRrqBmvHW6+5fRfpIx7A9jdsGY9QfcX+2H7LtadlQdR3bewJtWt7dzjPaPgy1oEDUmQEHI7RejrCsT2k2lGy+0+BPNsbu7pYZTZcNeyjSLnFgc5m7IaI8XDlmqWz3o1p3Wx3+IMQvXdlcYNiGJWtGgIcH2tanT3ahIza7fJy+PBegbEXeHbTW15tLeYTdYRtC7BcRp1q1Cyc2yxUdg8OrBwAayoCO8BkTpmYEmDMqf4Vstm7J3rVxQ2AvroUaJD39pc12VSyBqd0DLX3oOOxf0eYZjnpBs7HZ6qcPs8bwMYthVDOqDV7EuNuXE5d5lQTJjII9AexuD7T2+0tTGcMOIV7JtsKFEm4IG+94cS2h3z4R0HFbuwD7q+tNn8Vxu0u8IbsRToNZeV2ltK47W9YAxwcBugMe4SDMieELa2Fw+7p7S7aW1vY7OnZ+hiN5RsL3ErSi513eOc4UKDKr/FmATrAHUIN2x9FOy9e9tqNXZGm2k+q1rndlizYBIBMuG6MuJgcTkvJNh9jNn8S9Jm19hizapwjBhcvo0hWe3f3bhtJgc9jHvPjHhaSTHCV6zcXbLLAMBqbRW2E4BjNxb1H3Vp/gw3febXqNBkEBshrcv8AdoQuW9HmKssNttuNo7lwtrC1trWh6xZYa6w3qb7mi2W0W5tc5rHZSSSTCDuL70S7GWV/h2ztfALcG/rNqtrjELl1YObTmowVG25aGgEu3XObORO7OXm20no12bd6Ytm9kLSnUsbe8L/WXW93Urvc2HOY4OqUmtEgat3xM8l7jhNWneYC6vY0qFGrWrUGW9KtY9hUbbtP2T3UzTnc3qbd1zhEMMHgfGdv6O1WFel/ArzYzCru6bbBrrCh6u51v276TalxuNyDWk1N50boBJOUGA2sH9Fmz7NltoBU2Rxl1VzbdtKpXxnD6lZs1MxSeO5T0zLhLhkIzXmPpg2Ew3ZDCtnL7DaGJWlTEvWRVtb66oXJp9k5gBbUpANMh0xmQV7Ts/UwbELLFdl6eE7FG/tKNK+xKvQs96xo1zWbTayA77Q06b6kmY3nQNCuC/7QeD3dxgeH3uEUsEo7M4TWr022mH7jXWjq1Tu70OIqb7abXS0ZEuB0lB8/oQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAIQhAJWkjRIhBr4/j2KbQ35vcbv7m/uiA3tLioXGBoByHRQ3mK315h1hY3VzUq2lg17Lak492kHu3nAeZMrOQg39nNqcc2bbdtwDFbzDvW2hlY21UsLwJiSM8pOY5qtbY7itti7cVoYleU8UD+09bFZ3a73PemfqslCDbxjaLGMbbR/a2JXV4KDqj6Qq1SQx1R5e8gcCXEkn+y3LD0o7c2Nu2hbbVYw2i3wtNy50fGTC4hCDrj6Qtr/X7m9btNjFK7uQxtarSu3sNQMBDd7dImAckv/eVtvMja/aCdJ/aFWf/ANlyCEHS3W2u095WqVbzaLF69SrQNs91S8qOLqRMmmST4SdRoubJlIhBqvxvE34EzBn39y7CmVvWGWhqE021II3g3QGCfimYnil5iZtzfXDq3q9FtvRDjlTpt0a0aAZk+ZJWahAoV/FMSvcVuhc4lc1rquKbKQqVn7ztxjQ1ok8gAFnoQdDabWbQWeB1sFtcaxGjhNZpa+0ZcOFItOZG7MQZz5rOwvE77B76ne4VeXFld0/BWt6ppvbwMEQVnoQdJju2e020Fq22xvH8Tv7YHeFK5unvZPOCYlZ1xi+IXWGWeG3F5WqWFmXut7dz+5SLzLi0aAkrMQg7XDvSbtrh1oy3stqMXp27AA1nrLnBoiIEzAy0TP8AvG2w/a37VG0mKevFjabqgrnvNaHboI0MbzokZEkhcahB1OGbd7UYbil3iNltBidO+uwBcVzcOc6tGm8TMxJidJU2J+kLa7FCHXu0eKVIpPoQLhzR2b4LmwIkHdEjoFyCEGvgWPYrs/eet4JiF3YXO7udrbVTTcWnOCRqOivbQ7ZbSbS21K3x7HMRxChSdvsp3Fdz2h0RMExOcT1K5pCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCAQhCD/9k="

		var shot utility.Shot
		shot.ShotCode = *utility.GenerateCode(&shot.ShotNum)
		shot.Picture = pic
		shot.ShotNum = edls[i].ShotNum
		shot.ShotType = edls[i].ShotType
		shot.FromClipName = edls[i].FromClipName
		shot.SourceFile = edls[i].SourceFile
		shot.StartTime = edls[i].StartTime
		shot.EndTime = edls[i].EndTime
		shot.EdlCode = edlCode
		shot.EdlName = edlName
		shot.ShotFlag = "0"
		//通过SourceFile关联素材表material_name查信息
		stmt, err := mysqlUtility.DBConn.Prepare("SELECT material_code, library_code, width, height, timecode_framerate FROM material WHERE status = 0 AND project_code = ? AND material_name = ?")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		result := stmt.QueryRow(projectCode, shot.SourceFile)
		err = result.Scan(&shot.MaterialCode, &shot.LibraryCode, &shot.Width, &shot.Height, &shot.ShotFps)
		if err != nil {
			return nil, err
		}
		shots = append(shots, shot)
	}

	return shots, nil
}

func InsertMultipleShot(userCode string, projectCode string, shots []utility.Shot) error {
	tx, _ := mysqlUtility.DBConn.Begin()
	for i := 0; i < len(shots); i++ {
		stmt, err := tx.Prepare("INSERT INTO `shot`(shot_code, project_code, material_code, library_code, picture, shot_num, start_time, end_time, from_clip_name, source_file, shot_type, shot_name, shot_fps, width, height, shot_detail, shot_status, edl_code, edl_name, shot_flag, user_code, status, insert_datetime, update_datetime) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		s := shots[i]
		_, err = stmt.Exec(s.ShotCode, projectCode, s.MaterialCode, s.LibraryCode, s.Picture, s.ShotNum, s.StartTime, s.EndTime, s.FromClipName, s.SourceFile, s.ShotType, s.ShotName, s.ShotFps, s.Width, s.Height, s.ShotDetail, s.ShotStatus, s.EdlCode, s.EdlName, s.ShotFlag, userCode, s.Status)
		if err != nil {
			tx.Rollback()
			return err
		}
		stmt.Close()
	}
	tx.Commit()
	return nil
}

type shotOut struct {
	ShotCode   string
	ShotName   string
	ShotStatus string
	Picture    string
	ShotFlag   string
	SourcePath string
	DpxPath    string
	JpgPath    string
	MovPath    string
}

func QueryShots(projectCode *string) (*[]shotOut, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT a.shot_code, a.shot_name, a.shot_status, a.picture, a.shot_flag, IF(b.library_path LIKE '', 'N', 'Y') AS source_path, IF(b.dpx_path LIKE '', 'N', 'Y') AS dpx_path, IF(b.jpg_path LIKE '', 'N', 'Y') AS jpg_path, IF(b.mov_path LIKE '', 'N', 'Y') AS mov_path FROM shot a LEFT JOIN library b ON a.library_code = b.library_code AND a.status = b.status WHERE a.status = 0 AND a.project_code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var shots []shotOut
	for result.Next() {
		var so shotOut
		err = result.Scan(&so.ShotCode, &so.ShotName, &so.ShotStatus, &so.Picture, &so.ShotFlag, &so.SourcePath, &so.DpxPath, &so.JpgPath, &so.MovPath)
		if err != nil {
			return nil, err
		}
		shots = append(shots, so)
	}
	return &shots, err
}

func QueryShotByShotCode(shotCode *string) (*utility.Shot, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT shot_code, shot_name, width, height, shot_fps, start_time, end_time, shot_status, shot_detail, source_file, shot_type, shot_flag FROM `shot` WHERE status = 0 AND shot_code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var s utility.Shot
	result := stmt.QueryRow(shotCode)
	err = result.Scan(&s.ShotCode, &s.ShotName, &s.Width, &s.Height, &s.ShotFps, &s.StartTime, &s.EndTime, &s.ShotStatus, &s.ShotDetail, &s.SourceFile, &s.ShotType, &s.ShotFlag)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func UpdateShot(s *utility.Shot) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot SET shot_name = ?, width = ?, height = ?, shot_fps = ?, start_time = ?, end_time = ?, shot_detail = ?, user_code = ?, update_datetime = NOW() WHERE status = 0 AND shot_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.ShotName, s.Width, s.Height, s.ShotFps, s.StartTime, s.EndTime, s.ShotDetail, s.UserCode, s.ShotCode)
	if err != nil {
		return err
	}

	return nil
}

func AddSingleShot(s *utility.Shot) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO `shot`(shot_code, project_code, shot_type, shot_name, shot_fps, width, height, shot_detail, shot_flag, user_code, status, insert_datetime, update_datetime, material_code, library_code, picture, shot_num, start_time, end_time, from_clip_name, source_file, shot_status, edl_code, edl_name) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.ShotCode, s.ProjectCode, s.ShotType, s.ShotName, s.ShotFps, s.Width, s.Height, s.ShotDetail, s.ShotFlag, s.UserCode, s.Status, s.MaterialCode, s.LibraryCode, s.Picture, s.ShotNum, s.StartTime, s.EndTime, s.FromClipName, s.SourceFile, s.ShotStatus, s.EdlCode, s.EdlName)
	if err != nil {
		return err
	}
	return nil
}

func ModifyShotName(s *utility.Shot) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `shot` SET shot_name = ?, user_code = ?, update_datetime = NOW() WHERE status = 0 AND shot_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.ShotName, s.UserCode, s.ShotCode)
	if err != nil {
		return err
	}
	return nil
}

func FindFolderShots(code string, id string) (*[]shotOut, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT a.shot_code, a.shot_name, a.shot_status, a.picture, a.shot_flag, IF(b.library_path LIKE '', 'N', 'Y') AS source_path, IF(b.dpx_path LIKE '', 'N', 'Y') AS dpx_path, IF(b.jpg_path LIKE '', 'N', 'Y') AS jpg_path, IF(b.mov_path LIKE '', 'N', 'Y') AS mov_path FROM shot a LEFT JOIN library b ON a.library_code = b.library_code AND a.status = b.status WHERE a.status = 0 AND a.shot_code IN (SELECT shot_code from shot_folder_data where status = 0 AND project_code = ? AND folder_id = ?)")
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(code, id)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var shots []shotOut
	for result.Next() {
		var so shotOut
		err = result.Scan(&(so.ShotCode), &so.ShotName, &so.ShotStatus, &so.Picture, &so.ShotFlag, &so.SourcePath, &so.DpxPath, &so.JpgPath, &so.MovPath)
		if err != nil {
			return nil, err
		}
		shots = append(shots, so)
	}
	return &shots, err
}

func DeleteSingleShot(s *utility.Shot) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot SET status = 1, user_code = ?, update_datetime = NOW() WHERE status = 0 AND shot_flag = '1' AND shot_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(s.UserCode, s.ShotCode)
	if err != nil {
		return err
	}
	r, err := rs.RowsAffected()
	if r == 0 || err != nil {
		return errors.New("No data delete!")
	}
	return nil
}

// --------------------------------------------------------------

//cover old shot   start a transaction   code is ProjectCode
//func CoverMultipleShot(code string, edls []*utility.EdlShot) error {
//	length := len(edls)
//	tx, _ := mysqlUtility.DBConn.Begin()
//	shots, err := CopyEdlShot(tx, length, edls)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}

//	stmt2, err2 := tx.Prepare("UPDATE `shot` SET status=1") //全部的‘删除’
//	if err2 != nil {
//		stmt2.Close()
//		tx.Rollback()
//		return err2
//	}
//	_, err2 = stmt2.Exec()
//	if err2 != nil {
//		stmt2.Close()
//		tx.Rollback()
//		return err2
//	}
//	stmt2.Close()

//	for i := length - 1; i >= 0; i-- {
//		stmt, err := tx.Prepare("INSERT INTO  `shot`(shot_code,project_code,material_code,shot_num,start_time,end_time,clip_name,source_file,shot_type,shot_name,shot_fps,width,height,status) VALUE(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
//		if err != nil {
//			tx.Rollback()
//			return err
//		}
//		_, err = stmt.Exec(shots[i].ShotCode, shots[i].ProjectCode, shots[i].MaterialCode, code, shots[i].ShotNum, shots[i].StartTime, shots[i].EndTime, shots[i].FromClipName, shots[i].SourceFile,
//			shots[i].ShotType, shots[i].ShotName, shots[i].ShotFps, shots[i].Width, shots[i].Height, shots[i].Status)
//		if err != nil {
//			stmt.Close()
//			tx.Rollback()
//			return err
//		}
//		stmt.Close()
//	}
//	tx.Commit()
//	return nil
//}
