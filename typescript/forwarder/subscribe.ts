/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 5.28.2
 * source: forwarder/subscribe.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../google/protobuf/empty";
import * as dependency_2 from "./entity";
import * as pb_1 from "google-protobuf";
export class SubscribeResponse extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        fail_arr?: string[];
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("fail_arr" in data && data.fail_arr != undefined) {
                this.fail_arr = data.fail_arr;
            }
        }
    }
    get fail_arr() {
        return pb_1.Message.getFieldWithDefault(this, 1, []) as string[];
    }
    set fail_arr(value: string[]) {
        pb_1.Message.setField(this, 1, value);
    }
    static fromObject(data: {
        fail_arr?: string[];
    }): SubscribeResponse {
        const message = new SubscribeResponse({});
        if (data.fail_arr != null) {
            message.fail_arr = data.fail_arr;
        }
        return message;
    }
    toObject() {
        const data: {
            fail_arr?: string[];
        } = {};
        if (this.fail_arr != null) {
            data.fail_arr = this.fail_arr;
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.fail_arr.length)
            writer.writeRepeatedString(1, this.fail_arr);
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): SubscribeResponse {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new SubscribeResponse();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    pb_1.Message.addToRepeatedField(message, 1, reader.readString());
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): SubscribeResponse {
        return SubscribeResponse.deserialize(bytes);
    }
}
