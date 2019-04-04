'use strict';

exports.handler = (event, context, callback) => {
    const { request } = event.Records[0].cf;
	const { uri } = request;

	const inRootS3 = uri.match(/^\/(\w+)\/(app\.js|locales)/i)
	if(inRootS3) {
	    request.uri = uri.substr(inRootS3[1].length + 1);
	}
	
    callback(null, request);
};
