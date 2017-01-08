<!-- iframe里日历-->
<!DOCTYPE html>
<html>
<head>
<link rel='stylesheet' href='/static/css/fullcalendar.min.css' />
<script src='/static/js/jquery-2.1.3.min.js'></script>
<script src='/static/js/moment.min.js'></script>
<script src='/static/js/fullcalendar.min.js'></script>
<script src='/static/js/fullcalendar.zh-cn.js'></script>
<style>
	/*body {
		margin: 0;
		padding: 0;
		font-family: "Lucida Grande",Helvetica,Arial,Verdana,sans-serif;
		font-size: 14px;
	}*/

	#script-warning {
		display: none;
		background: #eee;
		border-bottom: 1px solid #ddd;
		padding: 0 10px;
		line-height: 40px;
		text-align: center;
		font-weight: bold;
		font-size: 12px;
		color: red;
	}

	#loading {
		display: none;
		position: absolute;
		top: 10px;
		right: 10px;
	}

	#calendar {
		max-width: 900px;
		margin: 40px auto;
		padding: 0 10px;
	}

	body {
		margin: 40px 10px;
		padding: 0;
		font-family: "Lucida Grande",Helvetica,Arial,Verdana,sans-serif;
		font-size: 14px;
	}

	/*#calendar {
		max-width: 900px;
		margin: 0 auto;
	}*/

</style>
</head>
<body>
<script type="text/javascript">
$(document).ready(function() {
    // page is now ready, initialize the calendar...
    $('#calendar').fullCalendar({
        // put your options and callbacks here
        	header: {
				left: 'prev,next today',
				center: 'title',
				right: 'month,agendaWeek,agendaDay,listMonth'
			},
			defaultDate: '2016-12-12',
			navLinks: true, // can click day/week names to navigate views
			editable: true,
			eventLimit: true, // allow "more" link when too many events
			businessHours: true, // display business hours
			selectable: true,
			selectHelper: true,
			select: function(start, end) {
				var title = prompt('Event Title:');
				var eventData;
				if (title) {
					eventData = {
						title: title,
						start: start,
						end: end,
						// Color: getRandomColor(),
						textColor: getRandomColor(),
						backgroundColor: getRandomColor(),
						// borderColor: getRandomColor(),
                		className: 'done',
					};
					$('#calendar').fullCalendar('renderEvent', eventData, true); // stick? = true
				}
				$('#calendar').fullCalendar('unselect');
			},
			editable: true,
			// events: {
			// 	url: '/admin/getcalendar',
			// 	error: function() {
			// 		$('#script-warning').show();
			// 	}
			// },
			loading: function(bool) {
				$('#loading').toggle(bool);
			},
			events: [
				{
					title: 'All Day Event',
					start: '2016-12-01'
				},
				{
					title: 'Long Event',
					start: '2016-12-07',
					end: '2016-12-10'
				},
				{
					id: 999,
					title: 'Repeating Event',
					start: '2016-12-09T16:00:00'
				},
				{
					id: 999,
					title: 'Repeating Event',
					start: '2016-12-16T16:00:00'
				},
				{
					title: 'Conference',
					start: '2016-12-11',
					end: '2016-12-13'
				},
				{
					title: 'Meeting',
					start: '2016-12-12T10:30:00',
					end: '2016-12-12T12:30:00'
				},
				{
					title: 'Lunch',
					start: '2016-12-12T12:00:00'
				},
				{
					title: 'Meeting',
					start: '2016-12-12T14:30:00'
				},
				{
					title: 'Happy Hour',
					start: '2016-12-12T17:30:00'
				},
				{
					title: 'Dinner',
					start: '2016-12-12T20:00:00'
				},
				{
					title: 'Birthday Party',
					start: '2016-12-13T07:00:00'
				},
				{
					title: 'Click for Google',
					url: 'http://google.com/',
					start: '2016-12-28'
				}
			],
			dayClick: function(date, jsEvent, view) {
        		// alert('Clicked on: ' + date.format());
        		// alert('Coordinates: ' + jsEvent.pageX + ',' + jsEvent.pageY);
        		// alert('Current view: ' + view.name);
        		// change the day's background color just for fun
        		// $(this).css('background-color', getRandomColor());
    		},
    		//单击事件项时触发 
        	eventClick: function(calEvent, jsEvent, view) {
        		// alert('Event: ' + calEvent.title);
        		// alert('Coordinates: ' + jsEvent.pageX + ',' + jsEvent.pageY);
        		// alert('View: ' + view.name);
        		// change the border color just for fun
        		$(this).css('border-color', getRandomColor());
    		},
    		// events: [
      //   		{
      //       		title: 'My Event',
      //       		start: '2016-12-01',
      //       		url: 'http://google.com/'
      //   		}
      //   	// other events here
    		// ],
    		// eventClick: function(event) {
      //   		if (event.url) {
      //       		window.open(event.url);
      //       		return false;
      //   		}
    		// }
    })
});

// $('#calendar').fullCalendar({
//     weekends: false // will hide Saturdays and Sundays
// });

// $('#calendar').fullCalendar({
//     dayClick: function() {
//         alert('a day has been clicked!');
//     }
// });

// $('#calendar').fullCalendar('next');
	function getRandomColor(){ 
		var c = '#'; 
		var cArray = ['0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F']; 
		for(var i = 0; i < 6;i++){ 
			var cIndex = Math.round(Math.random()*15); 
			c += cArray[cIndex]; 
		} 
			return c; 
		} 
</script>
<div class="col-lg-12">
	<div id='calendar'></div>
</div>
<!-- <script>
	$(document).ready(function() {
		$('#calendar').fullCalendar({
			header: {
				left: 'prev,next today',
				center: 'title',
				right: 'month,basicWeek,basicDay'
			},
			defaultDate: '2016-12-12',
			navLinks: true, // can click day/week names to navigate views
			editable: true,
			eventLimit: true, // allow "more" link when too many events
			events: [
				{
					title: 'All Day Event',
					start: '2016-12-01'
				},
				{
					title: 'Long Event',
					start: '2016-12-07',
					end: '2016-12-10'
				},
				{
					id: 999,
					title: 'Repeating Event',
					start: '2016-12-09T16:00:00'
				},
				{
					id: 999,
					title: 'Repeating Event',
					start: '2016-12-16T16:00:00'
				},
				{
					title: 'Conference',
					start: '2016-12-11',
					end: '2016-12-13'
				},
				{
					title: 'Meeting',
					start: '2016-12-12T10:30:00',
					end: '2016-12-12T12:30:00'
				},
				{
					title: 'Lunch',
					start: '2016-12-12T12:00:00'
				},
				{
					title: 'Meeting',
					start: '2016-12-12T14:30:00'
				},
				{
					title: 'Happy Hour',
					start: '2016-12-12T17:30:00'
				},
				{
					title: 'Dinner',
					start: '2016-12-12T20:00:00'
				},
				{
					title: 'Birthday Party',
					start: '2016-12-13T07:00:00'
				},
				{
					title: 'Click for Google',
					url: 'http://google.com/',
					start: '2016-12-28'
				}
			]
		});
		
	});
</script> -->

</body>
</html>