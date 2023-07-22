export namespace main {
	
	export class FileInfoJson {
	    name: string;
	    size: number;
	    // Go type: time
	    modeTime: any;
	    isDir: boolean;
	    Sys: any;
	
	    static createFrom(source: any = {}) {
	        return new FileInfoJson(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.size = source["size"];
	        this.modeTime = this.convertValues(source["modeTime"], null);
	        this.isDir = source["isDir"];
	        this.Sys = source["Sys"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DirEntryJson {
	    FileInfoJson: FileInfoJson;
	    name: string;
	    type: string;
	    isDir: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DirEntryJson(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.FileInfoJson = this.convertValues(source["FileInfoJson"], FileInfoJson);
	        this.name = source["name"];
	        this.type = source["type"];
	        this.isDir = source["isDir"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

