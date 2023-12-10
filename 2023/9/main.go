package main

func main() {
	inputs := []string{
		"0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45",
		"-3 8 35 82 152 255 432 803 1648 3531 7478 15221 29521 54584 96585 164316 269975 430114 666765 1008764 1493294\n12 27 56 116 234 455 852 1535 2668 4532 7730 13740 26223 53845 115961 253452 548456 1158885 2376712 4723336 9101236\n3 13 35 69 115 173 243 325 419 525 643 773 915 1069 1235 1413 1603 1805 2019 2245 2483\n6 30 78 166 318 581 1054 1938 3622 6826 12819 23717 42864 75377 129249 218254 369843 645270 1187135 2324466 4792495\n8 20 48 96 170 278 430 638 916 1280 1748 2340 3078 3986 5090 6418 8000 9868 12056 14600 17538\n10 28 46 54 38 -14 -105 -228 -366 -458 -249 1115 5938 19484 52314 123455 264058 520584 955564 1642484 2649198\n13 20 32 61 128 259 472 746 961 800 -382 -3677 -10602 -22661 -39996 -58340 -63147 -19452 144180 559145 1448580\n-2 11 33 62 111 222 479 1020 2048 3841 6761 11262 17897 27324 40311 57740 80610 110039 147265 193646 250659\n1 4 10 31 92 244 593 1357 2983 6386 13408 27644 55869 110482 213789 405871 759865 1413024 2631366 4944409 9421686\n3 7 12 15 13 19 90 375 1197 3191 7547 16486 34292 69668 141160 289463 604654 1284710 2757373 5926008 12650378\n17 30 60 120 223 382 610 920 1325 1838 2472 3240 4155 5230 6478 7912 9545 11390 13460 15768 18327\n21 33 59 118 238 466 892 1702 3296 6536 13233 27057 55185 111223 220292 427707 813463 1515843 2768959 4961016 8722648\n-9 -10 -11 -21 -46 -77 -68 116 817 2826 7846 19354 44214 95751 199690 405660 809383 1595322 3118827 6065589 11764030\n29 43 61 87 125 179 253 351 477 635 829 1063 1341 1667 2045 2479 2973 3531 4157 4855 5629\n22 37 63 123 266 580 1213 2410 4574 8359 14803 25509 42882 70430 113137 177916 274150 414329 614791 896575 1286394\n12 33 72 146 285 545 1033 1958 3733 7174 13896 27139 53562 107194 218027 450155 938626 1961319 4073624 8351424 16814387\n3 14 42 102 213 412 780 1472 2751 5056 9198 16886 31949 62863 127558 262041 535227 1073638 2101418 4003474 7422421\n24 52 106 197 336 534 802 1151 1592 2136 2794 3577 4496 5562 6786 8179 9752 11516 13482 15661 18064\n11 25 41 74 161 371 830 1783 3738 7777 16182 33616 69219 140113 276942 532255 993104 1801280 3191010 5571077 9719076\n10 20 40 93 214 455 896 1669 3010 5378 9731 18136 35020 69547 139835 280008 551410 1059688 1979874 3592055 6330706\n8 3 -1 1 20 81 251 706 1868 4663 10979 24439 51648 104125 201191 374152 672192 1170475 1981047 3267229 5262300\n20 30 32 36 68 185 518 1367 3392 7978 17902 38496 79587 158627 305671 571375 1040288 1854006 3253320 5657156 9811858\n24 34 58 112 215 389 659 1053 1602 2340 3304 4534 6073 7967 10265 13019 16284 20118 24582 29740 35659\n28 50 76 101 126 162 239 423 850 1803 3884 8373 17940 38029 79551 164159 334592 674792 1348388 2670725 5240402\n6 28 60 99 142 186 228 265 294 312 316 303 270 214 132 21 -122 -300 -516 -773 -1074\n15 35 83 170 313 542 915 1559 2760 5130 9888 19312 37464 71390 133214 243998 441141 792793 1425825 2579170 4703067\n16 33 67 118 181 243 279 249 112 -102 -88 1182 6303 20868 55712 130223 276992 548079 1023161 1819798 3106001\n5 11 17 23 29 35 41 47 53 59 65 71 77 83 89 95 101 107 113 119 125\n3 -4 -12 -13 21 153 496 1228 2607 4986 8828 14721 23393 35727 52776 75778 106171 145608 195972 259391 338253\n4 23 71 160 299 494 748 1061 1430 1849 2309 2798 3301 3800 4274 4699 5048 5291 5395 5324 5039\n21 26 31 34 29 5 -57 -186 -414 -747 -1096 -1145 -128 3518 12662 32193 69987 138180 254805 445856 747847\n9 31 62 110 205 426 947 2108 4517 9189 17728 32558 57209 96664 157773 249740 384689 578315 850626 1226782 1738037\n13 16 30 71 167 367 758 1487 2780 4951 8414 13777 22261 37052 66970 135448 301067 702266 1650886 3821103 8618663\n11 19 47 123 291 608 1152 2068 3705 6947 13927 29448 63630 136561 286058 580098 1136307 2152823 3959591 7110229 12558601\n5 9 20 43 83 145 234 355 513 713 960 1259 1615 2033 2518 3075 3709 4425 5228 6123 7115\n15 31 60 111 193 316 504 829 1484 2937 6249 13700 29968 64313 134701 275967 556906 1117762 2254755 4612208 9615988\n23 35 60 115 234 475 930 1747 3173 5638 9935 17627 31966 59936 116767 235958 491706 1049212 2276709 4990886 10982149\n12 25 59 136 293 584 1077 1847 2984 4666 7391 12519 23345 47007 97629 201208 402876 777303 1443155 2582682 4467685\n-3 9 47 135 311 630 1169 2037 3398 5525 8924 14615 24759 44017 82383 160854 322367 652253 1315537 2623555 5150795\n2 5 10 17 26 37 50 65 82 101 122 145 170 197 226 257 290 325 362 401 442\n-4 -6 -2 10 22 13 -44 -158 -278 -243 250 1613 4196 7874 11252 10307 -3760 -47740 -151034 -362068 -756966\n5 11 17 23 29 35 41 47 53 59 65 71 77 83 89 95 101 107 113 119 125\n25 34 55 113 244 511 1037 2063 4053 7898 15328 29738 57785 112332 217615 417901 791397 1471766 2680305 4772635 8304630\n6 6 12 35 96 243 592 1406 3233 7146 15184 31219 62738 124554 246458 488644 971932 1935218 3841428 7567307 14734088\n10 37 83 157 270 445 734 1249 2237 4278 8775 19051 42585 95222 208596 442525 904788 1781491 3382187 6205049 11028720\n10 12 23 53 130 321 769 1760 3857 8183 17007 34899 70913 142637 283750 558419 1089359 2113367 4091778 7929318 15403483\n16 24 40 78 158 316 630 1285 2710 5824 12423 25723 51042 96554 173977 298962 490828 771136 1160410 1672092 2302558\n-2 -3 5 33 93 198 362 600 928 1363 1923 2627 3495 4548 5808 7298 9042 11065 13393 16053 19073\n6 18 48 115 259 570 1240 2662 5613 11576 23286 45634 87138 162298 295300 525729 917200 1570126 2640220 4364781 7099349\n14 18 25 32 34 25 6 8 141 679 2196 5803 13639 29992 63835 134237 281144 585528 1204995 2433762 4799608\n5 14 40 111 280 650 1412 2896 5635 10442 18500 31465 51582 81814 125984 188930 276673 396598 557648 770531 1047940\n15 43 83 133 194 276 422 770 1682 3982 9362 21046 44873 91112 177601 335238 617425 1115657 1983773 3472911 5977044\n23 49 93 164 273 435 678 1069 1791 3344 6997 15687 35645 79128 168750 344034 670951 1255371 2261525 3936766 6644121\n5 23 52 93 155 261 454 803 1409 2411 3992 6385 9879 14825 21642 30823 42941 58655 78716 103973 135379\n-4 2 14 32 56 86 122 164 212 266 326 392 464 542 626 716 812 914 1022 1136 1256\n13 35 66 100 129 142 119 13 -293 -1078 -2972 -7341 -17061 -38018 -81835 -170381 -342207 -659383 -1208682 -2080734 -3289587\n10 26 56 111 202 340 536 801 1146 1582 2120 2771 3546 4456 5512 6725 8106 9666 11416 13367 15530\n3 8 35 94 193 337 521 724 932 1247 2186 5368 14983 40823 104359 250545 569949 1238742 2588377 5222894 10210209\n9 17 33 70 169 422 1006 2234 4635 9078 16959 30490 53194 90866 153573 259825 444969 777283 1387353 2519314 4616675\n1 12 33 62 105 185 369 829 1966 4660 10784 24266 53244 114309 240623 497177 1009363 2016864 3976363 7758442 15023863\n15 25 49 96 183 345 645 1184 2111 3633 6025 9640 14919 22401 32733 46680 65135 89129 119841 158608 206935\n2 19 61 146 309 628 1272 2593 5302 10792 21694 42769 82256 153868 279935 497170 869070 1514685 2675139 4861228 9164251\n6 19 49 104 197 347 593 1046 2032 4428 10364 24557 56662 125197 263927 532382 1034222 1952271 3621122 6685061 12445361\n21 46 95 179 315 538 931 1685 3197 6206 11965 22473 40878 72351 126066 219442 386537 693434 1264591 2325355 4267009\n11 17 22 26 29 31 32 32 31 29 26 22 17 11 4 -4 -13 -23 -34 -46 -59\n-5 -7 -8 -6 14 93 327 936 2420 5882 13645 30367 64997 134168 268066 520551 986473 1830895 3337518 5987256 10582936\n14 28 59 110 177 255 355 537 967 2021 4499 10104 22548 50100 111331 247641 549500 1209122 2621819 5573322 11573251\n12 23 44 82 144 246 429 785 1496 2889 5510 10220 18316 31680 52959 85779 134996 206987 309984 454454 653528\n12 29 53 84 129 208 360 664 1300 2685 5744 12450 26948 57953 123802 262705 552584 1147665 2343003 4683740 9143551\n7 -2 -16 -40 -83 -160 -290 -468 -574 -159 1989 8652 25607 64092 145107 307091 619903 1210398 2308644 4329594 8014103\n-3 -7 -18 -27 -6 110 460 1317 3187 6998 14509 29233 58493 117847 240240 494203 1018741 2088039 4228012 8420949 16456976\n8 12 23 54 125 263 502 883 1454 2270 3393 4892 6843 9329 12440 16273 20932 26528 33179 41010 50153\n7 16 46 125 306 689 1451 2888 5479 9986 17607 30217 50809 84468 140715 239051 421299 777264 1494792 2952113 5880136\n1 -2 -5 0 29 115 322 770 1687 3519 7149 14301 28235 54874 104544 194553 352885 623340 1072511 1799054 2945777\n12 31 61 111 201 362 636 1076 1746 2721 4087 5941 8391 11556 15566 20562 26696 34131 43041 53611 66037\n14 21 34 78 195 446 913 1701 2940 4787 7428 11080 15993 22452 30779 41335 54522 70785 90614 114546 143167\n-3 -4 -5 -6 -7 -8 -9 -10 -11 -12 -13 -14 -15 -16 -17 -18 -19 -20 -21 -22 -23\n16 22 31 43 68 144 367 942 2277 5165 11147 23252 47513 96021 192877 385325 763702 1495744 2883375 5453529 10098978\n12 21 32 43 57 90 193 496 1281 3090 6873 14180 27400 50049 87108 145411 234082 365019 553422 818361 1183379\n24 37 54 85 164 361 791 1620 3079 5523 9633 16999 31629 63566 137042 307905 699106 1569821 3444696 7346621 15201842\n11 35 74 147 295 588 1134 2108 3839 7031 13263 26026 52725 108314 221555 445311 872811 1662475 3074672 5524717 9657507\n8 15 25 47 109 274 668 1536 3359 7086 14565 29305 57788 111721 211954 395416 727506 1324121 2390113 4284603 7628287\n2 20 56 124 260 533 1057 2012 3686 6553 11404 19549 33108 55408 91501 148815 237946 373594 575640 870354 1291716\n0 -6 -1 38 160 454 1073 2277 4512 8569 15917 29389 54559 102470 195021 375567 729607 1424800 2788160 5456230 10672617\n23 50 87 130 175 218 255 282 295 290 263 210 127 10 -145 -342 -585 -878 -1225 -1630 -2097\n10 14 13 6 -8 -30 -61 -102 -154 -218 -295 -386 -492 -614 -753 -910 -1086 -1282 -1499 -1738 -2000\n-4 -11 -22 -25 2 85 255 582 1286 3003 7358 18147 43704 101487 226631 487270 1010920 2028249 3943258 7442391 13659526\n15 20 31 66 162 397 935 2100 4483 9090 17563 32578 58706 104455 187185 344660 661156 1323941 2740284 5770239 12173035\n1 1 1 4 11 21 31 36 29 1 -59 -164 -329 -571 -909 -1364 -1959 -2719 -3671 -4844 -6269\n7 4 15 67 206 516 1151 2391 4758 9267 17940 34778 67467 130189 248018 463504 846185 1505918 2611085 4412909 7277308\n25 49 87 139 205 285 379 487 609 745 895 1059 1237 1429 1635 1855 2089 2337 2599 2875 3165\n14 22 52 127 293 630 1259 2347 4119 6907 11310 18615 31746 57175 108455 212329 418740 816524 1557120 2889287 5208587\n12 9 6 3 0 -3 -6 -9 -12 -15 -18 -21 -24 -27 -30 -33 -36 -39 -42 -45 -48\n-10 -6 21 94 242 507 967 1786 3322 6371 12713 26270 55416 117368 246295 508125 1027594 2037916 3972401 7633582 14504290\n23 27 25 15 6 38 213 740 1997 4613 9573 18349 33060 56664 93185 147978 228035 342335 502241 721947 1018978\n9 8 10 25 77 204 455 884 1541 2460 3644 5047 6553 7952 8913 8954 7409 3392 -4242 -16939 -36491\n12 33 61 95 134 177 223 271 320 369 417 463 506 545 579 607 628 641 645 639 622\n-1 0 15 64 182 437 961 1994 3941 7442 13455 23352 39028 63023 98657 150178 222923 323492 459935 641952 881106\n-7 -12 -13 6 80 283 758 1757 3691 7190 13173 22928 38202 61301 95200 143663 211373 304072 428711 593610 808628\n4 -3 -19 -43 -71 -100 -124 -100 149 1149 4215 12343 32045 77166 176791 391314 844158 1782301 3687773 7476105 14838237\n14 35 74 133 209 294 376 445 514 669 1174 2715 7057 18873 50566 134004 346942 868595 2089930 4822005 10673209\n28 49 85 143 230 353 519 735 1008 1345 1753 2239 2810 3473 4235 5103 6084 7185 8413 9775 11278\n5 11 39 105 234 472 902 1663 2971 5141 8609 13953 21912 33402 49528 71591 101089 139711 189323 251945 329718\n-1 15 46 102 206 412 834 1687 3342 6407 11878 21473 38385 68886 125504 232902 438137 827695 1554614 2880152 5235864\n16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36\n-2 1 8 23 48 74 78 47 75 619 3065 10873 31803 82202 195310 437481 939876 1959741 3996514 8005971 15783282\n12 18 27 48 109 275 668 1497 3115 6129 11598 21363 38562 68392 119189 203906 342078 562372 905829 1429914 2213499\n-6 -2 21 91 263 628 1318 2507 4408 7266 11347 16923 24253 33560 45004 58651 74438 92134 111297 131227 150915\n5 10 22 63 166 375 745 1342 2243 3536 5320 7705 10812 14773 19731 25840 33265 42182 52778 65251 79810\n8 10 25 58 124 257 522 1050 2127 4385 9174 19247 39973 81414 161769 312909 589010 1079644 1929119 3364376 5734362\n-4 5 37 109 246 487 891 1543 2560 4097 6353 9577 14074 20211 28423 39219 53188 71005 93437 121349 155710\n8 23 66 146 274 484 866 1622 3173 6365 12845 25704 50513 96910 180931 328316 579062 993539 1660532 2707622 4314372\n15 24 47 101 219 461 942 1889 3738 7293 14004 26487 49506 91763 169019 309422 562865 1018841 1844115 3367746 6276313\n24 37 48 63 99 200 478 1197 2935 6893 15473 33322 69158 138923 271284 517458 969114 1791161 3281122 5974137 10823045\n6 31 74 140 242 421 778 1513 2978 5773 10953 20500 38406 73113 142846 286900 588888 1223735 2551772 5306858 10973880\n-1 0 3 19 77 243 656 1580 3481 7173 14148 27322 52601 101906 198601 386649 745283 1411526 2613529 4718421 8299181\n7 6 15 60 181 430 873 1619 2908 5297 10003 19529 38873 77994 156916 316071 636449 1277132 2542203 4995278 9648529\n-5 -12 -23 -38 -57 -80 -107 -138 -173 -212 -255 -302 -353 -408 -467 -530 -597 -668 -743 -822 -905\n13 24 39 78 189 465 1064 2227 4280 7599 12514 19126 27002 34684 38881 33085 5135 -67086 -223135 -532690 -1116397\n1 9 15 27 72 202 499 1079 2095 3739 6243 9879 14958 21828 30871 42499 57149 75277 97351 123843 155220\n6 9 15 20 19 3 -43 -140 -318 -624 -1101 -1591 -941 5412 31777 117414 360964 995315 2545498 6153517 14232263\n11 20 31 48 88 200 490 1154 2520 5102 9686 17522 30833 54180 97987 187244 382041 823848 1835153 4124588 9191574\n18 36 75 148 276 498 884 1551 2686 4584 7713 12822 21112 34494 55962 90113 143850 227308 355047 547560 833148\n1 -2 -5 -8 -11 -14 -17 -20 -23 -26 -29 -32 -35 -38 -41 -44 -47 -50 -53 -56 -59\n4 12 30 67 134 243 403 612 845 1040 1088 842 178 -824 -1429 854 13051 53877 171831 488950 1300771\n16 35 61 88 120 185 357 789 1756 3718 7464 14525 28299 56789 118630 255375 555241 1197591 2535234 5246790 10620275\n10 31 65 127 252 512 1040 2061 3930 7177 12559 21119 34252 53778 82022 121901 177018 251763 351421 482287 651788\n19 31 64 131 245 419 666 999 1431 1975 2644 3451 4409 5531 6830 8319 10011 11919 14056 16435 19069\n8 29 66 121 191 266 339 452 816 2057 5654 14649 34723 75746 153923 294672 536384 935229 1571186 2555489 4039695\n10 23 54 130 304 663 1334 2487 4335 7149 11354 17863 28953 50203 94312 188008 383760 778627 1543334 2965568 5512550\n3 -3 -12 -24 -39 -57 -78 -102 -129 -159 -192 -228 -267 -309 -354 -402 -453 -507 -564 -624 -687\n7 15 39 104 261 613 1365 2923 6091 12449 25048 49663 97086 187509 359353 687798 1324488 2582800 5119038 10310471 21026384\n12 15 32 66 121 219 439 992 2362 5571 12670 27636 58009 117906 233603 453802 868141 1639695 3060706 5646945 10298072\n4 -6 -20 -34 -43 -29 68 409 1369 3752 9219 21177 46630 99908 209804 432464 873340 1722587 3309714 6186423 11255537\n21 29 48 92 175 311 514 798 1177 1665 2276 3024 3923 4987 6230 7666 9309 11173 13272 15620 18231\n7 26 75 183 401 819 1592 2975 5367 9364 15821 25923 41265 63941 96642 142763 206519 293070 408655 560735 758145\n13 17 13 9 36 158 484 1194 2608 5351 10704 21309 42582 85604 173095 351591 715467 1452357 2925207 5814013 11350484\n2 12 32 74 157 307 557 947 1524 2342 3462 4952 6887 9349 12427 16217 20822 26352 32924 40662 49697\n17 30 52 102 224 516 1174 2562 5326 10575 20152 37018 65806 113762 192763 324215 550907 964104 1761400 3362618 6629320\n6 18 49 103 187 311 497 818 1503 3173 7321 17217 39499 86779 181608 362080 689327 1257755 2211892 3783590 6389642\n4 16 35 61 94 134 181 235 296 364 439 521 610 706 809 919 1036 1160 1291 1429 1574\n18 42 77 137 261 533 1111 2264 4414 8178 14403 24185 38861 59961 89105 127828 177314 238018 309153 388017 469133\n-7 -15 -16 14 108 310 695 1417 2793 5429 10406 19592 36256 66359 121211 222636 412401 768461 1431572 2647054 4827980\n17 39 67 98 127 145 140 114 140 493 1898 5946 15736 36807 78429 155326 289907 515083 877749 1443010 2299229\n10 15 32 82 205 482 1082 2345 4912 9913 19224 35804 64123 110692 184706 298811 470006 720691 1079872 1584534 2281193\n14 32 60 109 205 390 723 1298 2308 4194 7931 15529 30881 61192 119398 228262 427250 783884 1412086 2501117 4360133\n12 28 74 171 346 627 1042 1628 2458 3701 5742 9406 16352 29730 55226 102657 188320 338346 593362 1014821 1693422\n9 33 82 166 301 529 957 1831 3677 7568 15620 31890 63965 125740 242286 458515 854997 1577711 2895694 5315506 9811937\n-6 -1 14 47 126 319 759 1683 3510 7012 13682 26482 51270 99366 191931 367108 691218 1275725 2302190 4058033 6986622\n10 6 -5 -16 -15 24 174 636 1917 5201 13089 31008 69768 149991 309458 614830 1179708 2191616 3951231 6928056 11837747\n12 17 19 18 14 7 -3 -16 -32 -51 -73 -98 -126 -157 -191 -228 -268 -311 -357 -406 -458\n21 42 79 133 207 311 474 774 1396 2728 5513 11109 22021 43180 85188 172322 361107 780610 1714475 3758677 8109983\n19 40 81 155 293 563 1104 2182 4281 8253 15567 28718 51883 91942 160018 273731 460407 761534 1238813 1982213 3120505\n-6 -1 29 111 286 614 1182 2115 3590 5853 9239 14195 21306 31324 45200 64119 89538 123227 167313 224327 297254\n9 9 19 59 161 373 778 1555 3127 6463 13621 28626 58755 116229 220164 398378 688253 1135269 1787015 2679389 3810269\n-2 -1 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18\n28 40 60 107 223 487 1028 2048 3888 7218 13519 26174 52750 109544 230373 483202 1000939 2037155 4064417 7943548 15208324\n23 51 88 131 179 241 365 717 1756 4577 11548 27485 61846 132863 275304 554921 1095148 2126502 4079099 7759262 14693112\n7 16 30 50 89 190 457 1112 2596 5734 11994 23914 45890 85768 158135 290944 538237 1003364 1879366 3516242 6528815\n26 34 47 85 182 389 784 1500 2796 5227 10026 19913 40743 84810 177464 370417 768600 1584494 3248139 6629639 13487222\n14 13 3 -17 -42 -47 40 393 1412 4008 10230 24565 56497 125349 269180 560760 1135678 2240835 4316449 8131919 15006308\n27 45 78 151 302 590 1120 2090 3865 7093 12904 23281 41768 74790 134010 240344 430503 767237 1354826 2361803 4053410\n12 30 69 145 290 560 1042 1863 3214 5411 9029 15192 26258 47563 91865 190106 415769 936541 2118186 4725540 10303103\n-2 10 35 81 159 279 443 635 813 920 961 1273 3310 11699 39236 118346 326191 837605 2032938 4710820 10495116\n18 31 38 40 39 48 129 479 1601 4624 11873 27837 60740 124987 244834 459718 831780 1456221 2475248 4096494 6616933\n15 24 33 42 51 60 69 78 87 96 105 114 123 132 141 150 159 168 177 186 195\n-4 0 24 91 234 503 978 1796 3217 5790 10759 21012 43184 92062 199316 429934 911742 1886248 3793006 7405037 14039898\n16 34 78 159 288 482 777 1253 2076 3562 6268 11115 19548 33738 56831 93249 149048 232338 353770 527095 769800\n2 8 19 48 123 290 615 1190 2152 3727 6315 10649 18117 31478 56512 105757 206630 418270 864951 1799758 3721671\n-6 -10 -12 -2 38 144 384 871 1770 3294 5684 9168 13894 19832 26640 33489 38842 40182 33684 13826 -27066\n26 41 69 123 226 431 854 1728 3494 6953 13511 25557 47022 84175 146720 249266 413250 669401 1060841 1646927 2507946\n22 30 45 68 96 135 221 445 978 2092 4173 7722 13340 21693 33453 49211 69358 93930 122413 153504 184824\n7 18 45 98 190 352 656 1246 2377 4462 8127 14274 24152 39436 62314 95582 142747 208138 297025 415746 571842\n12 17 24 33 44 57 72 89 108 129 152 177 204 233 264 297 332 369 408 449 492\n-8 -5 6 23 56 137 327 719 1448 2755 5232 10528 23069 53827 128046 300485 685010 1512056 3241282 6784892 13953044\n15 39 88 174 320 584 1093 2092 4035 7781 15008 29024 56255 108890 209612 400332 758867 1431347 2694964 5077138 9575526\n20 41 75 137 252 458 817 1439 2529 4475 8008 14496 26526 49195 93241 182874 375062 805225 1790549 4052739 9177060\n8 7 6 24 102 318 811 1819 3744 7283 13731 25698 48723 94679 188588 381812 777139 1576007 3166374 6283137 12292741\n17 41 81 156 294 531 924 1593 2807 5126 9606 18068 33426 60064 104249 174567 282373 442255 672527 995788 1439614\n17 32 65 139 290 569 1052 1857 3159 5193 8254 12754 19490 30432 50607 92178 182963 384282 829153 1805853 3942912\n12 28 50 94 190 397 845 1819 3900 8178 16552 32132 59758 106651 183211 303977 488764 763992 1164222 1733914 2529422\n18 18 21 37 79 160 287 464 727 1245 2532 5826 13702 30997 66136 132959 253160 459460 799647 1341627 2179641\n5 8 11 14 17 20 23 26 29 32 35 38 41 44 47 50 53 56 59 62 65\n7 16 37 79 160 319 645 1351 2940 6531 14433 31070 64360 127617 241954 439009 763641 1276214 2054585 3197669 4836732\n2 -6 -3 22 83 207 452 935 1869 3616 6798 12588 23450 44843 88801 180935 373412 768062 1556259 3086058 5972853\n9 23 63 146 300 577 1080 2012 3752 6954 12650 22323 37926 61934 97915 153188 247647 438065 877257 1945628 4523676\n13 34 79 171 341 637 1143 2002 3441 5812 9696 16173 27449 48165 87914 165790 318229 612030 1165333 2180567 3995069\n23 31 39 47 55 63 71 79 87 95 103 111 119 127 135 143 151 159 167 175 183\n12 27 42 69 150 369 873 1912 3905 7544 13979 25214 45031 81118 149761 285849 563981 1143365 2364740 4961330 10515881\n13 21 29 37 45 53 61 69 77 85 93 101 109 117 125 133 141 149 157 165 173\n17 41 78 132 227 423 836 1675 3327 6558 12969 25967 52711 107842 220459 447060 895557 1769905 3452992 6660448 12726811\n23 44 77 136 248 457 837 1529 2824 5321 10196 19625 37411 69872 127054 224340 384533 640498 1038455 1642022 2537114\n0 7 30 95 242 525 1012 1785 2940 4587 6850 9867 13790 18785 25032 32725 42072 53295 66630 82327 100650\n15 28 57 126 268 523 940 1583 2533 3873 5655 7910 10940 16536 31555 76636 204838 537453 1326300 3051149 6552691\n4 2 -4 -13 -24 -36 -48 -59 -68 -74 -76 -73 -64 -48 -24 9 52 106 172 251 344\n13 24 43 80 154 310 652 1394 2931 5932 11457 21100 37160 62842 102490 161854 248393 371616 543463 778728 1095526\n6 22 53 110 218 433 863 1693 3214 5856 10225 17144 27698 43283 65659 97007 139990 197818 274317 374002 502154\n-8 -13 -15 -7 26 122 362 894 1966 3979 7569 13711 23818 39844 64673 104015 173514 320425 680803 1615087 4003952\n7 27 69 137 237 380 595 965 1697 3245 6535 13412 27573 56526 115640 236357 482563 982743 1992189 4011287 8005981\n5 16 39 74 121 180 251 334 429 536 655 786 929 1084 1251 1430 1621 1824 2039 2266 2505",
	}
	for _, input := range inputs {
		partOne(input)
		partTwo(input)
	}
}
