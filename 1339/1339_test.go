package main

import (
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	root := &TreeNode{
		Val: 1,
	}

	root.Left = &TreeNode{
		Val: 2,
	}

	root.Left.Left = &TreeNode{
		Val: 4,
	}
	root.Left.Right = &TreeNode{
		Val: 5,
	}

	root.Right = &TreeNode{
		Val: 3,
	}

	root.Right.Left = &TreeNode{
		Val: 6,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{root: root}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.root, 0, nil); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{args: []int{1, 2, 3}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.args...); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProduct(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{root: getExample1()}, want: 110},
		{name: "", args: args{root: getExample2()}, want: 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.root); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getExample1() *TreeNode {

	root := &TreeNode{
		Val: 1,
	}

	root.Left = &TreeNode{
		Val: 2,
	}

	root.Left.Left = &TreeNode{
		Val: 4,
	}
	root.Left.Right = &TreeNode{
		Val: 5,
	}

	root.Right = &TreeNode{
		Val: 3,
	}

	root.Right.Left = &TreeNode{
		Val: 6,
	}
	return root
}

func getExample2() *TreeNode {

	root := &TreeNode{
		Val: 1,
	}

	root.Right = &TreeNode{
		Val: 2,
	}

	root.Right.Left = &TreeNode{
		Val: 3,
	}
	root.Right.Right = &TreeNode{
		Val: 4,
	}

	root.Right.Right.Left = &TreeNode{
		Val: 5,
	}

	root.Right.Right.Right = &TreeNode{
		Val: 6,
	}
	return root
}

/*
//763478770
[3434,4223,2441,6764,5911,7094,1827,9223,3580,6615,8446,2770,5112,718,3292,4092,3269,377,7407,4515,4512,6098,282,2197,9833,5285,5841,9643,8708,500,1834,7466,1360,8075,9353,804,656,8645,2445,4648,1194,2185,7883,null,3282,2067,8329,4847,1363,1037,2829,3789,1321,8183,2392,8978,436,7776,2286,8635,587,4391,5075,7307,8431,2236,3588,null,null,6968,6324,null,2149,null,5868,7401,null,8175,7064,1404,8772,null,null,2259,3610,2455,7961,null,3397,8996,7112,1316,4197,8704,2391,227,4720,9266,3273,3503,null,5237,7905,1921,8540,1886,6681,4740,6134,8408,3442,null,3830,2786,5382,3499,4469,1260,1456,6568,746,8076,1665,null,5700,7959,209,2485,null,3253,6181,1080,8731,4829,7285,null,2136,3995,3153,4968,549,3290,627,7812,4406,254,8382,null,null,null,null,4246,5958,8358,1853,2260,3188,1963,9753,null,8976,null,1244,null,null,9473,8385,null,2370,3469,2059,9616,1238,3089,9857,null,3873,2465,1945,6202,7906,9853,2006,null,null,null,9707,4539,4815,3158,4493,5941,7648,7964,6534,9843,7333,409,5246,null,null,7899,null,2515,5559,2888,null,9258,9370,9767,null,4140,9421,5873,6398,1152,6895,9101,2652,6053,721,null,null,null,303,4819,1605,3110,5366,5363,6173,6188,7605,6414,null,1038,3255,870,798,217,1880,76,726,8296,4856,null,581,null,null,3247,261,7183,5597,6299,null,6591,null,2059,4841,3925,null,null,null,8182,3954,4280,5532,null,null,1900,1084,1704,459,4312,null,null,5370,6139,5517,1806,2949,3257,null,7820,7347,null,null,1224,1430,9544,1349,3583,5668,null,2646,3945,1422,9511,null,null,5411,8568,null,5311,7155,9720,1904,3772,null,3186,null,null,9768,2530,null,null,null,7092,null,7664,null,6784,2692,2031,9754,9864,4004,null,4518,4412,8741,5750,9149,null,6613,1859,null,784,2634,6172,4635,null,null,2675,null,null,718,5809,584,null,4540,1547,5399,null,2446,null,3829,null,null,null,8734,6278,8937,5207,2497,null,4524,502,9879,778,409,1892,null,7085,9924,2346,null,8991,null,null,9519,null,null,384,null,3435,5565,null,4053,7290,null,null,null,4105,2830,5815,8548,null,null,9641,8344,3660,null,9721,115,8064,3557,7344,6233,null,6000,9789,8861,3678,1170,1082,6525,6463,9355,null,null,null,1227,7890,9396,null,null,2978,7651,null,null,null,null,7948,3720,3218,1878,null,null,null,null,null,3308,null,null,null,null,null,null,null,201,4360,3255,null,null,2568,3229,null,3261,6406,null,5929,9516,null,null,null,9247,5988,2073,null,null,null,8689,null,null,null,null,5819,null,null,null,1352,null,9355,1350,null,null,9825,6797,1522,null,9138,null,null,1040,null,null,8289,6770,6913,7863,5710,6803,null,2297,null,null,null,9738,null,null,1004,null,null,null,null,4112,2947,913,8498,8879,null,null,8526,2609,3486,7499,null,null,9878,1833,309,5298,6989,3188,8401,7437,2560,831,null,4949,null,null,5723,8576,1991,null,7982,4985,4376,4170,5564,5641,null,null,1139,2392,null,null,null,4040,null,null,null,null,null,9909,null,null,2298,null,null,2294,496,9336,304,8230,null,6056,null,null,5382,759,7933,5960,3873,4548,6075,2686,7177,null,4091,null,767,null,4877,1756,1717,4537,null,null,null,null,2095,7406,870,null,null,null,1509,8381,null,3923,5682,667,4522,null,1869,7024,9323,1579,null,null,null,null,null,null,null,2951,3890,null,2131,null,8785,757,null,5912,2230,7837,4311,1439,6591,5966,null,2452,7591,2279,9438,8189,2293,8306,2934,3316,7320,4994,4027,null,6188,2554,null,9148,null,null,null,null,null,1252,5554,7583,1552,3862,null,8131,null,745,null,null,8880,16,7642,7181,null,null,2739,null,4819,null,null,null,null,null,1490,null,9057,null,null,null,1775,null,null,null,null,null,3482,null,7145,null,2374,9308,null,null,8204,null,null,null,928,null,null,null,null,null,3209,6076,null,null,9683,null,4879,null,7874,null,4504,69,null,null,null,1,362,null,1404,8888,1875,6483,6565,8395,2214,8420,4913,7525,7657,327,5459,2016,6299,5042,8558,876,4798,null,7907,4536,null,4508,null,null,814,6337,null,2782,3291,5829,null,7395,9107,9361,1260,10000,769,8659,null,null,9771,null,1160,5252,4088,3967,2779,7333,null,null,872,null,null,null,null,null,null,null,null,8543,null,null,9733,5123,4585,null,2880,9590,null,null,null,5920,null,5902,7266,null,9077,null,null,3794,null,2867,null,1822,null,null,3104,4782,9380,3245,null,8841,29,null,null,4115,4113,9770,4838,null,null,null,4573,null,null,null,5290,null,null,8254,null,null,5139,6,null,null,null,null,null,null,9066,null,4679,1378,1120,6088,2864,null,null,4555,null,null,4250,9565,7891,8402,3625,null,2381,null,null,null,null,null,980,null,5230,6426,5967,750,null,null,4537,null,null,null,null,null,3030,null,null,864,null,2600,null,8858,1373,2775,6504,null,5279,null,null,null,3933,9352,8292,1023,null,null,5354,3052,null,6710,null,2240,null,null,2165,null,null,null,2958,null,4988,null,5882,null,null,null,8686,3471,244,9758,96,null,3855,null,null,null,null,3210,null,8274,null,null,null,null,null,null,null,4513,null,null,5767,null,null,null,null,9506,null,null,null,null,null,null,null,null,null,6745,5337,null,3575,9091,null,null,null,9830,null,9738,7999,null,6138,null,9425,4203,3839,3824,null,9591,9675,null,6884,56,2355,5197,8046,4136,6357,null,8317,4117,null,1842,7985,6869,null,null,9812,8336,null,null,null,null,null,null,null,7511,1385,null,2986,null,8634,308,4008,7522,null,null,null,null,5879,null,null,2858,null,4856,5299,8482,5326,2624,858,4350,7321,521,8806,6941,3319,null,8915,1065,3453,6126,5694,null,null,null,null,9768,123,6371,null,null,2088,5298,782,9818,null,8197,1175,1146,1945,null,null,null,null,null,null,null,4079,null,null,null,3231,null,null,null,null,8841,8371,8971,null,1755,1891,null,null,null,null,3016,null,null,3136,null,678,null,7919,644,null,null,null,null,8769,null,3159,null,null,null,6179,null,null,null,null,null,2913,7271,null,733,3795,null,null,null,null,null,null,7164,2469,2744,5399,null,null,null,null,null,null,null,null,null,8888,7666,null,6396,1148,2642,null,4113,555,1669,null,null,null,null,3039,9576,7126,null,263,null,null,null,null,9748,null,null,9176,5235,222,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,6188,null,null,null,6012,8488,null,null,null,null,888,null,null,2384,null,6752,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,8105,1530,8038,null,null,null,9791,5982,4757,2791,9973,null,1030,null,null,null,5403,2796,null,null,4324,4362,1124,7488,4436,2556,5331,null,8660,1594,7398,8895,9023,null,null,577,5560,5592,null,3838,null,336,null,null,null,null,5396,null,5866,null,null,null,2308,null,null,9077,null,null,1270,null,1593,1990,null,null,6756,null,null,9650,null,null,5707,1106,null,5615,6076,825,null,null,6629,null,null,null,null,1502,4739,null,null,688,null,5134,4677,null,null,null,null,9885,1152,null,null,null,null,1619,8418,7699,8855,null,null,8402,9539,null,null,null,null,null,null,null,9468,null,null,null,null,null,5624,null,9683,null,null,null,2742,null,null,null,null,null,8847,7327,null,2759,2537,null,null,null,1981,null,null,4756,5394,1265,2611,null,4864,7675,null,null,null,null,null,213,null,null,null,151,null,null,null,7739,5814,130,null,null,null,481,1623,4669,null,8861,9953,5835,1593,4338,8037,2690,7015,6100,652,1497,null,null,null,null,null,null,6487,null,null,null,null,null,null,null,null,null,null,null,null,null,4692,2297,null,5542,null,null,null,null,9787,null,null,null,149,2236,4955,null,null,null,null,6703,null,5427,4698,2349,4578,9504,null,null,null,null,9921,473,null,null,null,null,null,null,null,null,null,null,2242,6092,4113,6763,null,null,8592,null,null,null,null,null,null,null,9095,null,null,null,9075,null,null,null,null,8514,null,null,5463,4782,2528,1399,null,3188,null,null,8314,null,null,null,null,null,null,3321,1832,8620,4797,9036,null,1243,null,null,null,7464,null,null,9351,null,null,null,null,null,7468,425,null,3747,null,null,null,7599,null,3724,null,742,9578,6532,2358,8082,340,9320,9391,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,5069,null,null,null,3669,null,598,629,null,null,null,null,null,null,null,null,null,2557,null,null,null,null,7496,6361,null,3693,null,1769,5807,4896,3384,2709,8316,3453,null,8683,6104,1462,null,3274,5488,null,9127,1830,null,null,null,4329,null,8761,null,null,null,null,1267,null,6006,2603,null,8455,null,null,312,null,5447,9453,null,7068,5531,4462,null,null,null,null,null,null,null,null,null,null,null,null,null,4103,null,null,3212,2659,null,null,null,null,1520,1032,null,null,null,null,null,null,1495,3098,709,9217,null,null,null,null,null,null,null,8529,4539,2454,8339,2077,6314,5155,7575,5197,7167,null,null,null,9379,null,null,null,9393,8408,null,8633,null,null,null,null,6241,8583,null,82,null,null,null,429,9391,6592,418,6105,5964,null,9746,8472,null,null,null,5501,null,3570,null,null,null,null,9490,null,null,null,null,null,null,null,null,null,null,null,3962,null,null,2299,7952,356,null,null,null,5180,8541,1860,null,null,null,null,4862,null,1825,null,3981,null,null,null,5022,null,null,null,null,null,6694,null,null,5648,2911,5099,null,null,null,null,null,null,null,null,8719,3130,null,null,null,null,null,null,null,null,9453,null,null,null,null,null,5808,null,null,null,null,null,null,null,null,9101,877,5009,1889,9232,3699,null,null,null,null,null,null,null,null,null,null,2971,6579,null,null,null,null,null,3152,null,null,8947,null,7526,null,null,8076,null,null,null,null,null,null,4827,4544,null,null,null,null,1373,null,null,null,null,null,null,null,null,9681,null,null,null,null,null,null,null,null,3068,null,4929,null,null,null,null,null,null,null,null,null,null,null,null,null,null,1447,null,171,null,4767,null,null,null,7197,null,null,3116,2056,5742,null,7142,7645,692,null,null,null,null,null,null,null,null,null,null,7244,null,null,null,null,9336,9541,1885,null,null,null,null,1095,null,null,null,null,null,null,null,8997,null,null,3604,4473,null,null,null,7284,298,null,null,null,null,null,null,9330,null,null,null,null,null,null,5501,null,null,null,null,2950,5479,null,null,null,null,null,null,null,null,null,8774,5861,null,2043]
*/
