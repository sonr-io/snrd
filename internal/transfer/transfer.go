package transfer

import (
	"bytes"
	"encoding/base64"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type OnProgress func(data float32)
type Transfer struct {
	// Inherited Properties
	mutex      sync.Mutex
	meta       *md.Metadata
	owner      *md.Peer
	onProgress OnProgress
	onComplete OnCompleted
	path       string

	// Builders
	stringsBuilder *strings.Builder
	bytesBuilder   *bytes.Buffer

	// Tracking
	currentSize int
	isBase64    bool
	interval    int
	totalChunks int
	totalSize   int
}

// ^ Check file type and use corresponding method ^ //
func (t *Transfer) addBuffer(curr int, buffer []byte) (bool, error) {
	// ** Lock/Unlock ** //
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// @ Unmarshal Bytes into Proto
	chunk := md.Chunk{}
	err := proto.Unmarshal(buffer, &chunk)
	if err != nil {
		return true, err
	}

	// @ Initialize Vars if First Chunk
	if curr == 0 {
		// Set Size
		t.totalSize = int(chunk.Total)

		// Check for base64
		if chunk.GetB64() != "" {
			// Calculate Tracking Data
			t.totalChunks = t.totalSize / B64ChunkSize
			t.interval = t.totalChunks / 100

			// Set Tracking Data
			t.isBase64 = true
		} else {
			// Set Tracking Data
			t.totalChunks = t.totalSize / BufferChunkSize
			t.isBase64 = false
		}
	}

	// @ Add Buffer by File Type
	if t.isBase64 {
		// Add Base64 Chunk to Buffer
		n, err := t.stringsBuilder.WriteString(chunk.B64)
		if err != nil {
			return true, err
		}

		// Update Tracking
		t.currentSize = t.currentSize + n
	} else {
		// Add ByteChunk to Buffer
		n, err := t.bytesBuilder.Write(chunk.Buffer)
		if err != nil {
			return true, err
		}

		// Update Tracking
		t.currentSize = t.currentSize + n
	}

	// @ Check Completed
	if t.currentSize < t.totalSize {
		// Check for Interval
		if curr%t.interval == 0 {
			// Send Callback
			t.onProgress(float32(t.currentSize) / float32(t.totalSize))
		}
		return false, nil
	} else {
		return true, nil
	}
}

// ^ Check file type and use corresponding method to save to Disk ^ //
func (t *Transfer) save() error {
	// @ Set File Bytes by Type
	if t.isBase64 {
		// Get Bytes from base64
		b64Bytes, err := base64.StdEncoding.DecodeString(t.stringsBuilder.String())
		if err != nil {
			log.Fatal("error:", err)
		}

		// Create File at Path
		f, err := os.Create(t.path)
		if err != nil {
			return err
		}
		defer f.Close()

		// Save Bytes from Base64
		if _, err := f.Write(b64Bytes); err != nil {
			return err
		}

		// Sync file
		if err := f.Sync(); err != nil {
			return err
		}

		// Create Metadata
		saved := &md.Metadata{
			Name:       t.meta.Name,
			Path:       t.path,
			Size:       int32(t.totalSize),
			Mime:       t.meta.Mime,
			Owner:      t.owner,
			LastOpened: int32(time.Now().Unix()),
		}

		// Convert Message to bytes
		bytes, err := proto.Marshal(saved)
		if err != nil {
			return err
		}

		// Send Complete Callback
		t.onComplete(true, bytes)
		return nil
	} else {
		// Create File at Path
		f, err := os.Create(t.path)
		if err != nil {
			return err
		}
		defer f.Close()

		// Save Bytes from Buffer
		if _, err := f.Write(t.bytesBuilder.Bytes()); err != nil {
			return err
		}

		// Sync file
		if err := f.Sync(); err != nil {
			return err
		}

		// Create Metadata
		saved := &md.Metadata{
			Name:       t.meta.Name,
			Path:       t.path,
			Size:       int32(t.totalSize),
			Mime:       t.meta.Mime,
			Owner:      t.owner,
			LastOpened: int32(time.Now().Unix()),
		}

		// Convert Message to bytes
		bytes, err := proto.Marshal(saved)
		if err != nil {
			return err
		}

		// Send Complete Callback
		t.onComplete(true, bytes)
		return nil
	}
}
