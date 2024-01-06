BOOL IsLittleEndian()
{
	int n = 1;
	if (*(char*)&n == 1) {
		TRUE;
	}
	return FALSE;
}

unsigned char* Prefix_Packet_Header(int bufferLength)
{
	unsigned char outputPacket[4] = {};
	unsigned char bigEndianMask[4] = { 24, 16, 8, 0 };
	BOOL isLittleEndian = IsLittleEndian();
	
	for (int i = 0; i < 4; i++)
	{
   		// Bitwise Operations
		if (isLittleEndian == TRUE) {
			// 0, 8, 16, 24
			outputPacket[i] = (unsigned char)((bufferLength >> (bigEndianMask[3 - i])) & 0xFF);
		}
		else {
			// 24, 16, 8, 0
			outputPacket[i] = (unsigned char)((bufferLength >> (bigEndianMask[i])) & 0xFF);
		}
	}
	return outputPacket;
}

/// You do not need to make a Bitwise Operations to return the Length of the packet in CPP, just receive it in an Int Variable, 

// Int Size In Memory IS 4 Bytes, and you just sent 4 Bytes over the network ; )
