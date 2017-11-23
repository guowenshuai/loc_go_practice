
	
	function getLoverTime() 
	{
		var loveDay = new Date();
		loveDay.setFullYear(2017,10,1)
		loveDay.setHours(0)
		loveDay.setMinutes(0)
		loveDay.setSeconds(0)
		
		
		var t = new Date() - loveDay
		t /= 1000
		var minute = 60 
		var hour = minute*60
		var day = 24*hour
		
		var days = Math.floor(t/day)
		days += 31
		t %= day
		var hours = Math.floor(t/hour)
		t %= hour
		var minutes = Math.floor(t/minute)
		t %= minute
		
		var timeString = days + " Days " + hours + " Hours " + minutes + " Minutes " + t + " Seconds"
		
		var s = document.getElementById("timeShow")
		s.innerHTML = timeString
		setTimeout(getLoverTime, 5*1000)
		
	}
	
