package seaweed

const resp = `
	[{
		"timestamp":1443592800,
		"localTimestamp":1443571200,
		"issueTimestamp":1443592800,
		"fadedRating":3,
		"solidRating":0,
		"swell":{
			"minBreakingHeight":5,
			"absMinBreakingHeight":4.88,
			"maxBreakingHeight":8,
			"absMaxBreakingHeight":7.63,
			"unit":"ft",
			"components":{
				"combined":{
					"height":7.5,
					"period":10,
					"direction":305.22,
					"compassDirection":"SE"
				},
				"primary":{
					"height":7.5,
					"period":10,
					"direction":309.5,
					"compassDirection":"SE"
				}
			}
		},
		"wind":{
			"speed":13,
			"direction":337,
			"compassDirection":"SSE",
			"chill":74,
			"gusts":27,
			"unit":"mph"
		},
		"condition":{
			"pressure":1008,
			"temperature":73,
			"weather":"22",
			"unitPressure":"mb",
			"unit":"f"
		},
		"charts":{
			"swell":"http:\/\/hist-2.msw.ms\/wave\/750\/20-1443592800-1.gif",
			"period":"http:\/\/hist-2.msw.ms\/wave\/750\/20-1443592800-2.gif",
			"wind":"http:\/\/hist-2.msw.ms\/gfs\/750\/20-1443592800-4.gif",
			"pressure":"http:\/\/hist-2.msw.ms\/gfs\/750\/20-1443592800-3.gif",
			"sst":"http:\/\/hist-2.msw.ms\/sst\/750\/20-1443592800-10.gif"
		}
	}]`
