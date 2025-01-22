export namespace api {
	
	export class ApiReturnValue {
	
	
	    static createFrom(source: any = {}) {
	        return new ApiReturnValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

