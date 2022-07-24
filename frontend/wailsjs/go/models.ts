export namespace models {
	
	export class EMDictionary {
	    id: number;
	    word: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new EMDictionary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.word = source["word"];
	        this.content = source["content"];
	    }
	}

}

