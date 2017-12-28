
let dataSet = [
 { label: "Game of Thrones", y: 0 , id:'game_of_thrones'},
 { label: "The Walking Death", y: 0, id:'the_walking_dead' },
 { label: "Stranger Things", y: 0 , id: 'stranger_things'},
 { label: "Mr. Robot", y: 0, id: 'mr_robot' },
]
const container = document.querySelector('#canvas');
const key = document.querySelector('#key').textContent;
const cluster = document.querySelector('#cluster').textContent;

if(container) {
 var chart = new CanvasJS.Chart("canvas",
  {
   animationEnabled: true,
   theme: "light2",
   title:{
   },
   data: [
    {
     type: "column",
     dataPoints: dataSet
    }
   ]
  });
 chart.render();

 // Just for dev purposes
 Pusher.logToConsole = true;

 var pusher = new Pusher(key, {
  encrypted: true,
  cluster: cluster
 });

 var channel = pusher.subscribe('tv-shows');
 channel.bind('vote-event', (data) => {
  dataSet = dataSet.map(x => {
   if(x.id == data.vote) {
    x.y += 10;
    return x
   } else {
    return x
   }
  });
  chart.render()
 });
}