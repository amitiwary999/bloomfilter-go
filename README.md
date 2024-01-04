Bloom filter algorithm implementation in golang.

## observation of experiment
1. 1 hash function
    1.
         1000 size of bf
		 100 item insert.
		  when we quired 80 non-existent item got 8 wrong info that item is present but it was not.
		  if 4,8, 16 wrong item quired 0 wrong info. 40 wrong item 5 wrong info.
	2.
	     1000 size of bf
		 1000 item insert.
		 80 non existent query 51 wrong return, 40 non existent return 21 wrong, 20 non existent return 14 wrong, 10 non existent return 5

	3.
	    10000 size of bf
		100 item insert
		40,80 non existent 1 wrong,
