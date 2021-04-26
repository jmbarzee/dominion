package ident

import (
	"github.com/jmbarzee/dominion/grpc"
)

type DomainRecord struct {
	// DomainIdentity holds the identifying information of the domain
	DomainIdentity

	// Services
	Services []ServiceIdentity
}

func NewDomainRecord(grpcDomainRecord *grpc.DomainRecord) (DomainRecord, error) {
	domainIdentity, err := NewDomainIdentity(grpcDomainRecord.DomainIdentity)
	if err != nil {
		return DomainRecord{}, err
	}

	serviceIdentityList, err := NewServiceIdentityList(grpcDomainRecord.GetServices())
	if err != nil {
		return DomainRecord{}, err
	}

	return DomainRecord{
		DomainIdentity: domainIdentity,
		Services:       serviceIdentityList,
	}, nil
}

func NewGRPCDomainRecord(domainRecord DomainRecord) *grpc.DomainRecord {
	return &grpc.DomainRecord{
		DomainIdentity: NewGRPCDomainIdentity(domainRecord.DomainIdentity),
		Services:       NewGRPCServiceIdentityList(domainRecord.Services),
	}
}

func NewDomainRecordList(grpcDomainRecords []*grpc.DomainRecord) ([]DomainRecord, error) {
	domainRecords := make([]DomainRecord, len(grpcDomainRecords))
	for i, grpcDomainRecord := range grpcDomainRecords {
		domainRecord, err := NewDomainRecord(grpcDomainRecord)
		if err != nil {
			return nil, err
		}
		domainRecords[i] = domainRecord
	}
	return domainRecords, nil
}

func NewGRPCDomainRecordList(domainRecords []DomainRecord) []*grpc.DomainRecord {
	grpcDomainRecords := make([]*grpc.DomainRecord, len(domainRecords))
	for i, domainRecord := range domainRecords {
		grpcDomainRecords[i] = NewGRPCDomainRecord(domainRecord)
	}
	return grpcDomainRecords
}
