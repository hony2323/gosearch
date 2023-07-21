export namespace main {
	
	export class DirEntryJson {
	    name: string;
	    info: string;
	    isDir: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DirEntryJson(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.info = source["info"];
	        this.isDir = source["isDir"];
	    }
	}

}

