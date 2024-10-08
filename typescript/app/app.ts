/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 5.28.2
 * source: app/app.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../forwarder/history";
import * as dependency_2 from "./../forwarder/realtime";
import * as dependency_3 from "./../forwarder/basic";
import * as dependency_4 from "./../forwarder/mq";
import * as pb_1 from "google-protobuf";
export enum PickListType {
    TYPE_ADD = 0,
    TYPE_REMOVE = 1
}
export class PickRealMap extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        pick_map?: Map<string, PickListType>;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("pick_map" in data && data.pick_map != undefined) {
                this.pick_map = data.pick_map;
            }
        }
        if (!this.pick_map)
            this.pick_map = new Map();
    }
    get pick_map() {
        return pb_1.Message.getField(this, 1) as any as Map<string, PickListType>;
    }
    set pick_map(value: Map<string, PickListType>) {
        pb_1.Message.setField(this, 1, value as any);
    }
    static fromObject(data: {
        pick_map?: {
            [key: string]: PickListType;
        };
    }): PickRealMap {
        const message = new PickRealMap({});
        if (typeof data.pick_map == "object") {
            message.pick_map = new Map(Object.entries(data.pick_map));
        }
        return message;
    }
    toObject() {
        const data: {
            pick_map?: {
                [key: string]: PickListType;
            };
        } = {};
        if (this.pick_map != null) {
            data.pick_map = (Object.fromEntries)(this.pick_map);
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        for (const [key, value] of this.pick_map) {
            writer.writeMessage(1, this.pick_map, () => {
                writer.writeString(1, key);
                writer.writeEnum(2, value);
            });
        }
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): PickRealMap {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new PickRealMap();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    reader.readMessage(message, () => pb_1.Map.deserializeBinary(message.pick_map as any, reader, reader.readString, reader.readEnum));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): PickRealMap {
        return PickRealMap.deserialize(bytes);
    }
}
export class PickFuture extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        code?: string;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("code" in data && data.code != undefined) {
                this.code = data.code;
            }
        }
    }
    get code() {
        return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
    }
    set code(value: string) {
        pb_1.Message.setField(this, 1, value);
    }
    static fromObject(data: {
        code?: string;
    }): PickFuture {
        const message = new PickFuture({});
        if (data.code != null) {
            message.code = data.code;
        }
        return message;
    }
    toObject() {
        const data: {
            code?: string;
        } = {};
        if (this.code != null) {
            data.code = this.code;
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.code.length)
            writer.writeString(1, this.code);
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): PickFuture {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new PickFuture();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    message.code = reader.readString();
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): PickFuture {
        return PickFuture.deserialize(bytes);
    }
}
export class WSMessage extends pb_1.Message {
    #one_of_decls: number[][] = [[1, 2, 3, 4, 5]];
    constructor(data?: any[] | ({} & (({
        future_detail?: dependency_3.FutureDetailMessage;
        future_tick?: never;
        history_kbar?: never;
        snapshot?: never;
        trade_index?: never;
    } | {
        future_detail?: never;
        future_tick?: dependency_4.FutureRealTimeTickMessage;
        history_kbar?: never;
        snapshot?: never;
        trade_index?: never;
    } | {
        future_detail?: never;
        future_tick?: never;
        history_kbar?: dependency_1.HistoryKbarResponse;
        snapshot?: never;
        trade_index?: never;
    } | {
        future_detail?: never;
        future_tick?: never;
        history_kbar?: never;
        snapshot?: dependency_2.SnapshotMessage;
        trade_index?: never;
    } | {
        future_detail?: never;
        future_tick?: never;
        history_kbar?: never;
        snapshot?: never;
        trade_index?: TradeIndex;
    })))) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("future_detail" in data && data.future_detail != undefined) {
                this.future_detail = data.future_detail;
            }
            if ("future_tick" in data && data.future_tick != undefined) {
                this.future_tick = data.future_tick;
            }
            if ("history_kbar" in data && data.history_kbar != undefined) {
                this.history_kbar = data.history_kbar;
            }
            if ("snapshot" in data && data.snapshot != undefined) {
                this.snapshot = data.snapshot;
            }
            if ("trade_index" in data && data.trade_index != undefined) {
                this.trade_index = data.trade_index;
            }
        }
    }
    get future_detail() {
        return pb_1.Message.getWrapperField(this, dependency_3.FutureDetailMessage, 1) as dependency_3.FutureDetailMessage;
    }
    set future_detail(value: dependency_3.FutureDetailMessage) {
        pb_1.Message.setOneofWrapperField(this, 1, this.#one_of_decls[0], value);
    }
    get has_future_detail() {
        return pb_1.Message.getField(this, 1) != null;
    }
    get future_tick() {
        return pb_1.Message.getWrapperField(this, dependency_4.FutureRealTimeTickMessage, 2) as dependency_4.FutureRealTimeTickMessage;
    }
    set future_tick(value: dependency_4.FutureRealTimeTickMessage) {
        pb_1.Message.setOneofWrapperField(this, 2, this.#one_of_decls[0], value);
    }
    get has_future_tick() {
        return pb_1.Message.getField(this, 2) != null;
    }
    get history_kbar() {
        return pb_1.Message.getWrapperField(this, dependency_1.HistoryKbarResponse, 3) as dependency_1.HistoryKbarResponse;
    }
    set history_kbar(value: dependency_1.HistoryKbarResponse) {
        pb_1.Message.setOneofWrapperField(this, 3, this.#one_of_decls[0], value);
    }
    get has_history_kbar() {
        return pb_1.Message.getField(this, 3) != null;
    }
    get snapshot() {
        return pb_1.Message.getWrapperField(this, dependency_2.SnapshotMessage, 4) as dependency_2.SnapshotMessage;
    }
    set snapshot(value: dependency_2.SnapshotMessage) {
        pb_1.Message.setOneofWrapperField(this, 4, this.#one_of_decls[0], value);
    }
    get has_snapshot() {
        return pb_1.Message.getField(this, 4) != null;
    }
    get trade_index() {
        return pb_1.Message.getWrapperField(this, TradeIndex, 5) as TradeIndex;
    }
    set trade_index(value: TradeIndex) {
        pb_1.Message.setOneofWrapperField(this, 5, this.#one_of_decls[0], value);
    }
    get has_trade_index() {
        return pb_1.Message.getField(this, 5) != null;
    }
    get data() {
        const cases: {
            [index: number]: "none" | "future_detail" | "future_tick" | "history_kbar" | "snapshot" | "trade_index";
        } = {
            0: "none",
            1: "future_detail",
            2: "future_tick",
            3: "history_kbar",
            4: "snapshot",
            5: "trade_index"
        };
        return cases[pb_1.Message.computeOneofCase(this, [1, 2, 3, 4, 5])];
    }
    static fromObject(data: {
        future_detail?: ReturnType<typeof dependency_3.FutureDetailMessage.prototype.toObject>;
        future_tick?: ReturnType<typeof dependency_4.FutureRealTimeTickMessage.prototype.toObject>;
        history_kbar?: ReturnType<typeof dependency_1.HistoryKbarResponse.prototype.toObject>;
        snapshot?: ReturnType<typeof dependency_2.SnapshotMessage.prototype.toObject>;
        trade_index?: ReturnType<typeof TradeIndex.prototype.toObject>;
    }): WSMessage {
        const message = new WSMessage({});
        if (data.future_detail != null) {
            message.future_detail = dependency_3.FutureDetailMessage.fromObject(data.future_detail);
        }
        if (data.future_tick != null) {
            message.future_tick = dependency_4.FutureRealTimeTickMessage.fromObject(data.future_tick);
        }
        if (data.history_kbar != null) {
            message.history_kbar = dependency_1.HistoryKbarResponse.fromObject(data.history_kbar);
        }
        if (data.snapshot != null) {
            message.snapshot = dependency_2.SnapshotMessage.fromObject(data.snapshot);
        }
        if (data.trade_index != null) {
            message.trade_index = TradeIndex.fromObject(data.trade_index);
        }
        return message;
    }
    toObject() {
        const data: {
            future_detail?: ReturnType<typeof dependency_3.FutureDetailMessage.prototype.toObject>;
            future_tick?: ReturnType<typeof dependency_4.FutureRealTimeTickMessage.prototype.toObject>;
            history_kbar?: ReturnType<typeof dependency_1.HistoryKbarResponse.prototype.toObject>;
            snapshot?: ReturnType<typeof dependency_2.SnapshotMessage.prototype.toObject>;
            trade_index?: ReturnType<typeof TradeIndex.prototype.toObject>;
        } = {};
        if (this.future_detail != null) {
            data.future_detail = this.future_detail.toObject();
        }
        if (this.future_tick != null) {
            data.future_tick = this.future_tick.toObject();
        }
        if (this.history_kbar != null) {
            data.history_kbar = this.history_kbar.toObject();
        }
        if (this.snapshot != null) {
            data.snapshot = this.snapshot.toObject();
        }
        if (this.trade_index != null) {
            data.trade_index = this.trade_index.toObject();
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.has_future_detail)
            writer.writeMessage(1, this.future_detail, () => this.future_detail.serialize(writer));
        if (this.has_future_tick)
            writer.writeMessage(2, this.future_tick, () => this.future_tick.serialize(writer));
        if (this.has_history_kbar)
            writer.writeMessage(3, this.history_kbar, () => this.history_kbar.serialize(writer));
        if (this.has_snapshot)
            writer.writeMessage(4, this.snapshot, () => this.snapshot.serialize(writer));
        if (this.has_trade_index)
            writer.writeMessage(5, this.trade_index, () => this.trade_index.serialize(writer));
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): WSMessage {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new WSMessage();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    reader.readMessage(message.future_detail, () => message.future_detail = dependency_3.FutureDetailMessage.deserialize(reader));
                    break;
                case 2:
                    reader.readMessage(message.future_tick, () => message.future_tick = dependency_4.FutureRealTimeTickMessage.deserialize(reader));
                    break;
                case 3:
                    reader.readMessage(message.history_kbar, () => message.history_kbar = dependency_1.HistoryKbarResponse.deserialize(reader));
                    break;
                case 4:
                    reader.readMessage(message.snapshot, () => message.snapshot = dependency_2.SnapshotMessage.deserialize(reader));
                    break;
                case 5:
                    reader.readMessage(message.trade_index, () => message.trade_index = TradeIndex.deserialize(reader));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): WSMessage {
        return WSMessage.deserialize(bytes);
    }
}
export class TradeIndex extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        tse?: IndexStatus;
        otc?: IndexStatus;
        nasdaq?: IndexStatus;
        nf?: IndexStatus;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("tse" in data && data.tse != undefined) {
                this.tse = data.tse;
            }
            if ("otc" in data && data.otc != undefined) {
                this.otc = data.otc;
            }
            if ("nasdaq" in data && data.nasdaq != undefined) {
                this.nasdaq = data.nasdaq;
            }
            if ("nf" in data && data.nf != undefined) {
                this.nf = data.nf;
            }
        }
    }
    get tse() {
        return pb_1.Message.getWrapperField(this, IndexStatus, 1) as IndexStatus;
    }
    set tse(value: IndexStatus) {
        pb_1.Message.setWrapperField(this, 1, value);
    }
    get has_tse() {
        return pb_1.Message.getField(this, 1) != null;
    }
    get otc() {
        return pb_1.Message.getWrapperField(this, IndexStatus, 2) as IndexStatus;
    }
    set otc(value: IndexStatus) {
        pb_1.Message.setWrapperField(this, 2, value);
    }
    get has_otc() {
        return pb_1.Message.getField(this, 2) != null;
    }
    get nasdaq() {
        return pb_1.Message.getWrapperField(this, IndexStatus, 3) as IndexStatus;
    }
    set nasdaq(value: IndexStatus) {
        pb_1.Message.setWrapperField(this, 3, value);
    }
    get has_nasdaq() {
        return pb_1.Message.getField(this, 3) != null;
    }
    get nf() {
        return pb_1.Message.getWrapperField(this, IndexStatus, 4) as IndexStatus;
    }
    set nf(value: IndexStatus) {
        pb_1.Message.setWrapperField(this, 4, value);
    }
    get has_nf() {
        return pb_1.Message.getField(this, 4) != null;
    }
    static fromObject(data: {
        tse?: ReturnType<typeof IndexStatus.prototype.toObject>;
        otc?: ReturnType<typeof IndexStatus.prototype.toObject>;
        nasdaq?: ReturnType<typeof IndexStatus.prototype.toObject>;
        nf?: ReturnType<typeof IndexStatus.prototype.toObject>;
    }): TradeIndex {
        const message = new TradeIndex({});
        if (data.tse != null) {
            message.tse = IndexStatus.fromObject(data.tse);
        }
        if (data.otc != null) {
            message.otc = IndexStatus.fromObject(data.otc);
        }
        if (data.nasdaq != null) {
            message.nasdaq = IndexStatus.fromObject(data.nasdaq);
        }
        if (data.nf != null) {
            message.nf = IndexStatus.fromObject(data.nf);
        }
        return message;
    }
    toObject() {
        const data: {
            tse?: ReturnType<typeof IndexStatus.prototype.toObject>;
            otc?: ReturnType<typeof IndexStatus.prototype.toObject>;
            nasdaq?: ReturnType<typeof IndexStatus.prototype.toObject>;
            nf?: ReturnType<typeof IndexStatus.prototype.toObject>;
        } = {};
        if (this.tse != null) {
            data.tse = this.tse.toObject();
        }
        if (this.otc != null) {
            data.otc = this.otc.toObject();
        }
        if (this.nasdaq != null) {
            data.nasdaq = this.nasdaq.toObject();
        }
        if (this.nf != null) {
            data.nf = this.nf.toObject();
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.has_tse)
            writer.writeMessage(1, this.tse, () => this.tse.serialize(writer));
        if (this.has_otc)
            writer.writeMessage(2, this.otc, () => this.otc.serialize(writer));
        if (this.has_nasdaq)
            writer.writeMessage(3, this.nasdaq, () => this.nasdaq.serialize(writer));
        if (this.has_nf)
            writer.writeMessage(4, this.nf, () => this.nf.serialize(writer));
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): TradeIndex {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new TradeIndex();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    reader.readMessage(message.tse, () => message.tse = IndexStatus.deserialize(reader));
                    break;
                case 2:
                    reader.readMessage(message.otc, () => message.otc = IndexStatus.deserialize(reader));
                    break;
                case 3:
                    reader.readMessage(message.nasdaq, () => message.nasdaq = IndexStatus.deserialize(reader));
                    break;
                case 4:
                    reader.readMessage(message.nf, () => message.nf = IndexStatus.deserialize(reader));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): TradeIndex {
        return TradeIndex.deserialize(bytes);
    }
}
export class IndexStatus extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        break_count?: number;
        price_chg?: number;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("break_count" in data && data.break_count != undefined) {
                this.break_count = data.break_count;
            }
            if ("price_chg" in data && data.price_chg != undefined) {
                this.price_chg = data.price_chg;
            }
        }
    }
    get break_count() {
        return pb_1.Message.getFieldWithDefault(this, 1, 0) as number;
    }
    set break_count(value: number) {
        pb_1.Message.setField(this, 1, value);
    }
    get price_chg() {
        return pb_1.Message.getFieldWithDefault(this, 2, 0) as number;
    }
    set price_chg(value: number) {
        pb_1.Message.setField(this, 2, value);
    }
    static fromObject(data: {
        break_count?: number;
        price_chg?: number;
    }): IndexStatus {
        const message = new IndexStatus({});
        if (data.break_count != null) {
            message.break_count = data.break_count;
        }
        if (data.price_chg != null) {
            message.price_chg = data.price_chg;
        }
        return message;
    }
    toObject() {
        const data: {
            break_count?: number;
            price_chg?: number;
        } = {};
        if (this.break_count != null) {
            data.break_count = this.break_count;
        }
        if (this.price_chg != null) {
            data.price_chg = this.price_chg;
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.break_count != 0)
            writer.writeInt64(1, this.break_count);
        if (this.price_chg != 0)
            writer.writeDouble(2, this.price_chg);
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): IndexStatus {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new IndexStatus();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    message.break_count = reader.readInt64();
                    break;
                case 2:
                    message.price_chg = reader.readDouble();
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): IndexStatus {
        return IndexStatus.deserialize(bytes);
    }
}
