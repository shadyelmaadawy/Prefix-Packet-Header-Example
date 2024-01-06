import Foundation

// MARK: NOTE
// MARK: Apple products define Integer ( Int ) as Int64 and use it as default data type due to the use of x64 Architecture In OS,
// MARK: And do not use x32 Architecture, so we are In ultra-wide operating system,
// MARK: The size will be 8 Bytes instead of 4, if you use Int32 Data Type size will be 4 as ordinary
// MARK: No need to check for little and big-endian

func Prefix_Packet_Header(bufferLength: Int) -> Data {
    var packetSize = packetLength
    return Data.init(
        bytes: &packetSize,
        count: MemoryLayout<Int>.size
    )
}


func Read_PacketLenght(packetBuffer: Data) -> Int {
    return packetBuffer.withUnsafeBytes { unsafePointer in
        return unsafePointer.load(as: Int.self)
    }
}

var packetLength = 128
let packetHeader = Prefix_Packet_Header(
    bufferLength: packetLength
)

print(packetHeader)
let packetSize = Read_PacketLenght(
    packetBuffer: packetHeader
)
print(packetSize)
