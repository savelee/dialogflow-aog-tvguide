const TVGUIDE_WEBSERVICE = 'BACKEND_URL/channel';
const { WebhookClient } = require('dialogflow-fulfillment');
const rp = require('request-promise');
const moment = require('moment');

process.env.DEBUG = 'dialogflow:debug'; // enables lib debugging statements

exports.tvguide = function(request, response){
  var agent = new WebhookClient({ request, response });
  
  let spokenText = "";
  let channelInput = request.body.queryResult.parameters.channel;
  let requestedTime = request.body.queryResult.parameters.time;
  let url = `${TVGUIDE_WEBSERVICE}/${channelInput}`;
  
  if (requestedTime) {
    // console.log(requestedTime); //2019-06-07T20:00:00+02:00
    let time = moment(requestedTime).add(2,'h').format('HH:mm:ss');
    url = `${TVGUIDE_WEBSERVICE}/${channelInput}/${time}`;
  }

  var options = {
      uri: url,
      json: true
  };
   
    // request promise calls an URL and returns the JSON response.
    rp(options)
      .then(function(tvresults) {
        // the JSON response, will need to be formatted in 'spoken' text strings.
        spokenText = speak(tvresults);
      })
      .catch(function (err) {
        console.error(err);
        spokenText = "Oh no! I can't retrieve tv information right now.";
      })
      .finally(function(){
        // kick start the Dialogflow app
        // based on an intent match, execute
        // the handler to return a spokenText string
        let intentMap = new Map();
        intentMap.set('Test Intent', testHandler);
        intentMap.set('Channel Intent', channelHandler);
        
        agent.handleRequest(intentMap);
      });

  function testHandler(agent) {
    agent.add(`Testing testing!`);
  }
  function channelHandler(agent) {
    // return this to Dialogflow / the Assistant
    agent.add(spokenText);
  }
};

/**
 * Return a text string to be spoken out by the Google Assistant
 * @param {object} JSON tv results
 */
var speak = function(tvresults) {
    let s = "";
    if(tvresults['Listings'][0]) {
        let channelName = getChannelName(tvresults['ID']);
        let currentlyPlayingTime = getSpokenTime(tvresults['Listings'][0]['Time']);
        let laterPlayingTime = getSpokenTime(tvresults['Listings'][1]['Time']);
        s = `On ${channelName} since ${currentlyPlayingTime}, ${tvresults['Listings'][0]['Title']} is playing.
        Afterwards at ${laterPlayingTime}, ${tvresults['Listings'][1]['Title']} will start.`
    }

    return s;
}

/**
 * Get ChannelName by Chanel Id number
 * @param {string} channelId representing a number
 * @return {string} channelName
 */
var getChannelName = function(channelId){
    let channels = new Map();
    channels.set(1, 'Net 1');
    channels.set(2, 'Net 2');
    channels.set(3, 'Net 3');
    channels.set(4, 'RTL 4');
    channels.set(5, 'Movie Channel');
    channels.set(6, 'Sports Channel');
    channels.set(7, 'Comedy Central');
    channels.set(8, 'Cartoon Network');
    channels.set(9, 'National Geographic');
    channels.set(10, 'MTV');

    return channels.get(channelId);
}

/**
 * Return a natural spoken time
 * @param {string} time in 'HH:mm:ss' format
 * @returns {string} spoken time (like 8 30 pm i.s.o. 20:00:00)
 */
var getSpokenTime = function(time){
    let datetime = moment(time, 'HH:mm:ss');
    let min = moment(datetime).format('m');
    let hour = moment(datetime).format('h');
    let partOfTheDay = moment(datetime).format('a');

    if (min == '0') {
        min = '';
    }

    return `${hour} ${min} ${partOfTheDay}`;
};