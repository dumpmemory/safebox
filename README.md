# 🔐 SafeBox
[![GoDoc][1]][2] [![MIT licensed][3]][4]

[1]: https://godoc.org/github.com/xtaci/safebox?status.svg
[2]: https://pkg.go.dev/github.com/xtaci/safebox
[3]: https://img.shields.io/badge/license-MIT-blue.svg
[4]: LICENSE

An unified key management system to make life easier. 

The main goal of safebox is to make key backup easier with single main key to derive the reset. You only need to backup **ONE** file about **256KB** to your removable disk, such as floppy disk/thumb drive/mo disk(magneto optical)/dvd-ram, etc...

# Features

1. **Unlimited keys** can be derived with a single main key, but we still suggest one master key for 16384 derived keys.
2. **Multi-source entropy**, entropy comes from key strokes, system entropy(/dev/urandom), startup time, process pid, etc...
3. **Plugable exporters** to adapt to different scenario, such as, blockchain, secure shell.

# Safebox Can Derive Keys For:

1. SSH
2. Ethereum
3. Bitcoin
4. Atom
5. Band
6. Persistence
7. Kava
8. Akash
9. Filecoin
10. NEM
11. Tron

and more plugable export plugin keeping coming...


# Recommendations

1. Install on Openbsd
2. Backup your master key file on **Removable & Reliable storages**, such as [DVD-RAM](https://en.wikipedia.org/wiki/DVD-RAM), [MO](https://en.wikipedia.org/wiki/Magneto-optical_drive). 
 
Common Storage Lifetime Table:

| Storage | LifeTime | Cold storage|
|------|------|------|
|Magneto Optical Disk| estimated 50 years | Yes |
|DVD-RAM|estimated 30 years| Yes |
|Floppy disk | 10-20 years | Yes |
|Flash(SD Card/USB Thumb/SSD)|5-10 years or more (depends on write cycles)| Yes |
|Hard Disk | 3-5 years| **No** |
|Recordable CD/DVD| 2-5 years | Yes|


# TUI

Safebox is designed with a **retro-style** text-based user interface(tui), so a box such as Raspberry Pi will be able to act as key mangement box for offline storage. And keys can be obtain via text-based **QR-Code**.

![image](https://github.com/user-attachments/assets/cdfb6c13-30f6-4892-ac38-53ae090614d6)
![image](https://github.com/user-attachments/assets/431d773d-b597-4dd0-a7aa-5b7c45a738cb)
![1731123361910](https://github.com/user-attachments/assets/d9cf031d-8ec0-4708-b91b-04bcb4ff9a2f)
![image](https://github.com/user-attachments/assets/be031eb0-2fa7-4b0e-8558-d21667a58f9e)




![image](https://user-images.githubusercontent.com/2346725/116670086-e595f080-a9d1-11eb-92b1-b5724b5e764e.png)


# Status 

Beta
