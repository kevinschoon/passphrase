# Passphrase

Generate a [diceware-style](http://world.std.com/~reinhold/diceware.html) random passphrase based on
new wordlists [published](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) by EFF.

![xkcd](http://imgs.xkcd.com/comics/password_strength.png)


## Usage

    passphrase
    correcthorsebatterystaple
    passphrase --roles 4 --camelCase --specialChar --digit
    correctHorseBatteryStaple!1
    passphrase --verbose
    Die[0/5] rolled: 3
    Die[1/5] rolled: 3
    Die[2/5] rolled: 4
    Die[3/5] rolled: 1
    Die[4/5] rolled: 4
    Dice rolled: 33414 (hangout)
    Die[0/5] rolled: 2
    Die[1/5] rolled: 3
    Die[2/5] rolled: 4
    Die[3/5] rolled: 1
    Die[4/5] rolled: 3
    Dice rolled: 23413 (discolor)
    Die[0/5] rolled: 1
    Die[1/5] rolled: 4
    Die[2/5] rolled: 3
    Die[3/5] rolled: 4
    Die[4/5] rolled: 3
    Dice rolled: 14343 (bulk)
    Die[0/5] rolled: 3
    Die[1/5] rolled: 1
    Die[2/5] rolled: 3
    Die[3/5] rolled: 5
    Die[4/5] rolled: 4
    Dice rolled: 31354 (foster)
    Passphrase entropy:  51.69925001442312
    hangoutDiscolorBulkFoster

