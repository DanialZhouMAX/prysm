package eth

import (
	"github.com/prysmaticlabs/prysm/v5/encoding/bytesutil"
)

// CopyETH1Data copies the provided eth1data object.
func CopyETH1Data(data *Eth1Data) *Eth1Data {
	if data == nil {
		return nil
	}
	return &Eth1Data{
		DepositRoot:  bytesutil.SafeCopyBytes(data.DepositRoot),
		DepositCount: data.DepositCount,
		BlockHash:    bytesutil.SafeCopyBytes(data.BlockHash),
	}
}

// CopyPendingAttestationSlice copies the provided slice of pending attestation objects.
func CopyPendingAttestationSlice(input []*PendingAttestation) []*PendingAttestation {
	if input == nil {
		return nil
	}

	res := make([]*PendingAttestation, len(input))
	for i := 0; i < len(res); i++ {
		res[i] = CopyPendingAttestation(input[i])
	}
	return res
}

// CopyPendingAttestation copies the provided pending attestation object.
func CopyPendingAttestation(att *PendingAttestation) *PendingAttestation {
	if att == nil {
		return nil
	}
	data := CopyAttestationData(att.Data)
	return &PendingAttestation{
		AggregationBits: bytesutil.SafeCopyBytes(att.AggregationBits),
		Data:            data,
		InclusionDelay:  att.InclusionDelay,
		ProposerIndex:   att.ProposerIndex,
	}
}

// CopyAttestation copies the provided attestation object.
func CopyAttestation(att *Attestation) *Attestation {
	if att == nil {
		return nil
	}
	return &Attestation{
		AggregationBits: bytesutil.SafeCopyBytes(att.AggregationBits),
		Data:            CopyAttestationData(att.Data),
		Signature:       bytesutil.SafeCopyBytes(att.Signature),
	}
}

// CopyAttestationElectra copies the provided attestation object.
func CopyAttestationElectra(att *AttestationElectra) *AttestationElectra {
	if att == nil {
		return nil
	}
	return &AttestationElectra{
		AggregationBits: bytesutil.SafeCopyBytes(att.AggregationBits),
		CommitteeBits:   bytesutil.SafeCopyBytes(att.CommitteeBits),
		Data:            CopyAttestationData(att.Data),
		Signature:       bytesutil.SafeCopyBytes(att.Signature),
	}
}

// CopyAttestationData copies the provided AttestationData object.
func CopyAttestationData(attData *AttestationData) *AttestationData {
	if attData == nil {
		return nil
	}
	return &AttestationData{
		Slot:            attData.Slot,
		CommitteeIndex:  attData.CommitteeIndex,
		BeaconBlockRoot: bytesutil.SafeCopyBytes(attData.BeaconBlockRoot),
		Source:          CopyCheckpoint(attData.Source),
		Target:          CopyCheckpoint(attData.Target),
	}
}

// CopyCheckpoint copies the provided checkpoint.
func CopyCheckpoint(cp *Checkpoint) *Checkpoint {
	if cp == nil {
		return nil
	}
	return &Checkpoint{
		Epoch: cp.Epoch,
		Root:  bytesutil.SafeCopyBytes(cp.Root),
	}
}

// CopySignedBeaconBlock copies the provided SignedBeaconBlock.
func CopySignedBeaconBlock(sigBlock *SignedBeaconBlock) *SignedBeaconBlock {
	if sigBlock == nil {
		return nil
	}
	return &SignedBeaconBlock{
		Block:     CopyBeaconBlock(sigBlock.Block),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBeaconBlock copies the provided BeaconBlock.
func CopyBeaconBlock(block *BeaconBlock) *BeaconBlock {
	if block == nil {
		return nil
	}
	return &BeaconBlock{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBeaconBlockBody(block.Body),
	}
}

// CopyBeaconBlockBody copies the provided BeaconBlockBody.
func CopyBeaconBlockBody(body *BeaconBlockBody) *BeaconBlockBody {
	if body == nil {
		return nil
	}
	return &BeaconBlockBody{
		RandaoReveal:      bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:          CopyETH1Data(body.Eth1Data),
		Graffiti:          bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings: CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings: CopyAttesterSlashings(body.AttesterSlashings),
		Attestations:      CopyAttestations(body.Attestations),
		Deposits:          CopyDeposits(body.Deposits),
		VoluntaryExits:    CopySignedVoluntaryExits(body.VoluntaryExits),
	}
}

// CopySignedBeaconBlockAltair copies the provided SignedBeaconBlock.
func CopySignedBeaconBlockAltair(sigBlock *SignedBeaconBlockAltair) *SignedBeaconBlockAltair {
	if sigBlock == nil {
		return nil
	}
	return &SignedBeaconBlockAltair{
		Block:     CopyBeaconBlockAltair(sigBlock.Block),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBeaconBlockAltair copies the provided BeaconBlock.
func CopyBeaconBlockAltair(block *BeaconBlockAltair) *BeaconBlockAltair {
	if block == nil {
		return nil
	}
	return &BeaconBlockAltair{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBeaconBlockBodyAltair(block.Body),
	}
}

// CopyBeaconBlockBodyAltair copies the provided BeaconBlockBody.
func CopyBeaconBlockBodyAltair(body *BeaconBlockBodyAltair) *BeaconBlockBodyAltair {
	if body == nil {
		return nil
	}
	return &BeaconBlockBodyAltair{
		RandaoReveal:      bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:          CopyETH1Data(body.Eth1Data),
		Graffiti:          bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings: CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings: CopyAttesterSlashings(body.AttesterSlashings),
		Attestations:      CopyAttestations(body.Attestations),
		Deposits:          CopyDeposits(body.Deposits),
		VoluntaryExits:    CopySignedVoluntaryExits(body.VoluntaryExits),
		SyncAggregate:     CopySyncAggregate(body.SyncAggregate),
	}
}

// CopyProposerSlashings copies the provided ProposerSlashing array.
func CopyProposerSlashings(slashings []*ProposerSlashing) []*ProposerSlashing {
	if slashings == nil {
		return nil
	}
	newSlashings := make([]*ProposerSlashing, len(slashings))
	for i, att := range slashings {
		newSlashings[i] = CopyProposerSlashing(att)
	}
	return newSlashings
}

// CopyProposerSlashing copies the provided ProposerSlashing.
func CopyProposerSlashing(slashing *ProposerSlashing) *ProposerSlashing {
	if slashing == nil {
		return nil
	}
	return &ProposerSlashing{
		Header_1: CopySignedBeaconBlockHeader(slashing.Header_1),
		Header_2: CopySignedBeaconBlockHeader(slashing.Header_2),
	}
}

// CopySignedBeaconBlockHeader copies the provided SignedBeaconBlockHeader.
func CopySignedBeaconBlockHeader(header *SignedBeaconBlockHeader) *SignedBeaconBlockHeader {
	if header == nil {
		return nil
	}
	return &SignedBeaconBlockHeader{
		Header:    CopyBeaconBlockHeader(header.Header),
		Signature: bytesutil.SafeCopyBytes(header.Signature),
	}
}

// CopyBeaconBlockHeader copies the provided BeaconBlockHeader.
func CopyBeaconBlockHeader(header *BeaconBlockHeader) *BeaconBlockHeader {
	if header == nil {
		return nil
	}
	parentRoot := bytesutil.SafeCopyBytes(header.ParentRoot)
	stateRoot := bytesutil.SafeCopyBytes(header.StateRoot)
	bodyRoot := bytesutil.SafeCopyBytes(header.BodyRoot)
	return &BeaconBlockHeader{
		Slot:          header.Slot,
		ProposerIndex: header.ProposerIndex,
		ParentRoot:    parentRoot,
		StateRoot:     stateRoot,
		BodyRoot:      bodyRoot,
	}
}

// CopyAttesterSlashings copies the provided AttesterSlashings array.
func CopyAttesterSlashings(slashings []*AttesterSlashing) []*AttesterSlashing {
	if slashings == nil {
		return nil
	}
	newSlashings := make([]*AttesterSlashing, len(slashings))
	for i, slashing := range slashings {
		newSlashings[i] = &AttesterSlashing{
			Attestation_1: CopyIndexedAttestation(slashing.Attestation_1),
			Attestation_2: CopyIndexedAttestation(slashing.Attestation_2),
		}
	}
	return newSlashings
}

// CopyAttesterSlashingsElectra copies the provided AttesterSlashings array.
func CopyAttesterSlashingsElectra(slashings []*AttesterSlashingElectra) []*AttesterSlashingElectra {
	if slashings == nil {
		return nil
	}
	newSlashings := make([]*AttesterSlashingElectra, len(slashings))
	for i, slashing := range slashings {
		newSlashings[i] = &AttesterSlashingElectra{
			Attestation_1: CopyIndexedAttestationElectra(slashing.Attestation_1),
			Attestation_2: CopyIndexedAttestationElectra(slashing.Attestation_2),
		}
	}
	return newSlashings
}

// CopyIndexedAttestation copies the provided IndexedAttestation.
func CopyIndexedAttestation(indexedAtt *IndexedAttestation) *IndexedAttestation {
	var indices []uint64
	if indexedAtt == nil {
		return nil
	} else if indexedAtt.AttestingIndices != nil {
		indices = make([]uint64, len(indexedAtt.AttestingIndices))
		copy(indices, indexedAtt.AttestingIndices)
	}
	return &IndexedAttestation{
		AttestingIndices: indices,
		Data:             CopyAttestationData(indexedAtt.Data),
		Signature:        bytesutil.SafeCopyBytes(indexedAtt.Signature),
	}
}

// CopyIndexedAttestationElectra copies the provided IndexedAttestation.
func CopyIndexedAttestationElectra(indexedAtt *IndexedAttestationElectra) *IndexedAttestationElectra {
	var indices []uint64
	if indexedAtt == nil {
		return nil
	} else if indexedAtt.AttestingIndices != nil {
		indices = make([]uint64, len(indexedAtt.AttestingIndices))
		copy(indices, indexedAtt.AttestingIndices)
	}
	return &IndexedAttestationElectra{
		AttestingIndices: indices,
		Data:             CopyAttestationData(indexedAtt.Data),
		Signature:        bytesutil.SafeCopyBytes(indexedAtt.Signature),
	}
}

// CopyAttestations copies the provided Attestation array.
func CopyAttestations(attestations []*Attestation) []*Attestation {
	if attestations == nil {
		return nil
	}
	newAttestations := make([]*Attestation, len(attestations))
	for i, att := range attestations {
		newAttestations[i] = CopyAttestation(att)
	}
	return newAttestations
}

// CopyAttestations copies the provided AttestationElectra array.
func CopyAttestationsElectra(attestations []*AttestationElectra) []*AttestationElectra {
	if attestations == nil {
		return nil
	}
	newAttestations := make([]*AttestationElectra, len(attestations))
	for i, att := range attestations {
		newAttestations[i] = CopyAttestationElectra(att)
	}
	return newAttestations
}

// CopyDeposits copies the provided deposit array.
func CopyDeposits(deposits []*Deposit) []*Deposit {
	if deposits == nil {
		return nil
	}
	newDeposits := make([]*Deposit, len(deposits))
	for i, dep := range deposits {
		newDeposits[i] = CopyDeposit(dep)
	}
	return newDeposits
}

// CopyDeposit copies the provided deposit.
func CopyDeposit(deposit *Deposit) *Deposit {
	if deposit == nil {
		return nil
	}
	return &Deposit{
		Proof: bytesutil.SafeCopy2dBytes(deposit.Proof),
		Data:  CopyDepositData(deposit.Data),
	}
}

// CopyDepositData copies the provided deposit data.
func CopyDepositData(depData *Deposit_Data) *Deposit_Data {
	if depData == nil {
		return nil
	}
	return &Deposit_Data{
		PublicKey:             bytesutil.SafeCopyBytes(depData.PublicKey),
		WithdrawalCredentials: bytesutil.SafeCopyBytes(depData.WithdrawalCredentials),
		Amount:                depData.Amount,
		Signature:             bytesutil.SafeCopyBytes(depData.Signature),
	}
}

// CopySignedVoluntaryExits copies the provided SignedVoluntaryExits array.
func CopySignedVoluntaryExits(exits []*SignedVoluntaryExit) []*SignedVoluntaryExit {
	if exits == nil {
		return nil
	}
	newExits := make([]*SignedVoluntaryExit, len(exits))
	for i, exit := range exits {
		newExits[i] = CopySignedVoluntaryExit(exit)
	}
	return newExits
}

// CopySignedVoluntaryExit copies the provided SignedVoluntaryExit.
func CopySignedVoluntaryExit(exit *SignedVoluntaryExit) *SignedVoluntaryExit {
	if exit == nil {
		return nil
	}
	return &SignedVoluntaryExit{
		Exit: &VoluntaryExit{
			Epoch:          exit.Exit.Epoch,
			ValidatorIndex: exit.Exit.ValidatorIndex,
		},
		Signature: bytesutil.SafeCopyBytes(exit.Signature),
	}
}

// CopyValidator copies the provided validator.
func CopyValidator(val *Validator) *Validator {
	pubKey := make([]byte, len(val.PublicKey))
	copy(pubKey, val.PublicKey)
	withdrawalCreds := make([]byte, len(val.WithdrawalCredentials))
	copy(withdrawalCreds, val.WithdrawalCredentials)
	return &Validator{
		PublicKey:                  pubKey,
		WithdrawalCredentials:      withdrawalCreds,
		EffectiveBalance:           val.EffectiveBalance,
		Slashed:                    val.Slashed,
		ActivationEligibilityEpoch: val.ActivationEligibilityEpoch,
		ActivationEpoch:            val.ActivationEpoch,
		ExitEpoch:                  val.ExitEpoch,
		WithdrawableEpoch:          val.WithdrawableEpoch,
	}
}

// CopySyncCommitteeMessage copies the provided sync committee message object.
func CopySyncCommitteeMessage(s *SyncCommitteeMessage) *SyncCommitteeMessage {
	if s == nil {
		return nil
	}
	return &SyncCommitteeMessage{
		Slot:           s.Slot,
		BlockRoot:      bytesutil.SafeCopyBytes(s.BlockRoot),
		ValidatorIndex: s.ValidatorIndex,
		Signature:      bytesutil.SafeCopyBytes(s.Signature),
	}
}

// CopySyncCommitteeContribution copies the provided sync committee contribution object.
func CopySyncCommitteeContribution(c *SyncCommitteeContribution) *SyncCommitteeContribution {
	if c == nil {
		return nil
	}
	return &SyncCommitteeContribution{
		Slot:              c.Slot,
		BlockRoot:         bytesutil.SafeCopyBytes(c.BlockRoot),
		SubcommitteeIndex: c.SubcommitteeIndex,
		AggregationBits:   bytesutil.SafeCopyBytes(c.AggregationBits),
		Signature:         bytesutil.SafeCopyBytes(c.Signature),
	}
}

// CopySyncAggregate copies the provided sync aggregate object.
func CopySyncAggregate(a *SyncAggregate) *SyncAggregate {
	if a == nil {
		return nil
	}
	return &SyncAggregate{
		SyncCommitteeBits:      bytesutil.SafeCopyBytes(a.SyncCommitteeBits),
		SyncCommitteeSignature: bytesutil.SafeCopyBytes(a.SyncCommitteeSignature),
	}
}

// CopySignedBeaconBlockBellatrix copies the provided SignedBeaconBlockBellatrix.
func CopySignedBeaconBlockBellatrix(sigBlock *SignedBeaconBlockBellatrix) *SignedBeaconBlockBellatrix {
	if sigBlock == nil {
		return nil
	}
	return &SignedBeaconBlockBellatrix{
		Block:     CopyBeaconBlockBellatrix(sigBlock.Block),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBeaconBlockBellatrix copies the provided BeaconBlockBellatrix.
func CopyBeaconBlockBellatrix(block *BeaconBlockBellatrix) *BeaconBlockBellatrix {
	if block == nil {
		return nil
	}
	return &BeaconBlockBellatrix{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBeaconBlockBodyBellatrix(block.Body),
	}
}

// CopyBeaconBlockBodyBellatrix copies the provided BeaconBlockBodyBellatrix.
func CopyBeaconBlockBodyBellatrix(body *BeaconBlockBodyBellatrix) *BeaconBlockBodyBellatrix {
	if body == nil {
		return nil
	}
	return &BeaconBlockBodyBellatrix{
		RandaoReveal:      bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:          CopyETH1Data(body.Eth1Data),
		Graffiti:          bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings: CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings: CopyAttesterSlashings(body.AttesterSlashings),
		Attestations:      CopyAttestations(body.Attestations),
		Deposits:          CopyDeposits(body.Deposits),
		VoluntaryExits:    CopySignedVoluntaryExits(body.VoluntaryExits),
		SyncAggregate:     CopySyncAggregate(body.SyncAggregate),
		ExecutionPayload:  body.ExecutionPayload.Copy(),
	}
}

// CopySignedBeaconBlockCapella copies the provided SignedBeaconBlockCapella.
func CopySignedBeaconBlockCapella(sigBlock *SignedBeaconBlockCapella) *SignedBeaconBlockCapella {
	if sigBlock == nil {
		return nil
	}
	return &SignedBeaconBlockCapella{
		Block:     CopyBeaconBlockCapella(sigBlock.Block),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBeaconBlockCapella copies the provided BeaconBlockCapella.
func CopyBeaconBlockCapella(block *BeaconBlockCapella) *BeaconBlockCapella {
	if block == nil {
		return nil
	}
	return &BeaconBlockCapella{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBeaconBlockBodyCapella(block.Body),
	}
}

// CopyBeaconBlockBodyCapella copies the provided BeaconBlockBodyCapella.
func CopyBeaconBlockBodyCapella(body *BeaconBlockBodyCapella) *BeaconBlockBodyCapella {
	if body == nil {
		return nil
	}
	return &BeaconBlockBodyCapella{
		RandaoReveal:          bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:              CopyETH1Data(body.Eth1Data),
		Graffiti:              bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings:     CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings:     CopyAttesterSlashings(body.AttesterSlashings),
		Attestations:          CopyAttestations(body.Attestations),
		Deposits:              CopyDeposits(body.Deposits),
		VoluntaryExits:        CopySignedVoluntaryExits(body.VoluntaryExits),
		SyncAggregate:         CopySyncAggregate(body.SyncAggregate),
		ExecutionPayload:      body.ExecutionPayload.Copy(),
		BlsToExecutionChanges: CopyBLSToExecutionChanges(body.BlsToExecutionChanges),
	}
}

// CopySignedBlindedBeaconBlockCapella copies the provided SignedBlindedBeaconBlockCapella.
func CopySignedBlindedBeaconBlockCapella(sigBlock *SignedBlindedBeaconBlockCapella) *SignedBlindedBeaconBlockCapella {
	if sigBlock == nil {
		return nil
	}
	return &SignedBlindedBeaconBlockCapella{
		Block:     CopyBlindedBeaconBlockCapella(sigBlock.Block),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBlindedBeaconBlockCapella copies the provided BlindedBeaconBlockCapella.
func CopyBlindedBeaconBlockCapella(block *BlindedBeaconBlockCapella) *BlindedBeaconBlockCapella {
	if block == nil {
		return nil
	}
	return &BlindedBeaconBlockCapella{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBlindedBeaconBlockBodyCapella(block.Body),
	}
}

// CopyBlindedBeaconBlockBodyCapella copies the provided BlindedBeaconBlockBodyCapella.
func CopyBlindedBeaconBlockBodyCapella(body *BlindedBeaconBlockBodyCapella) *BlindedBeaconBlockBodyCapella {
	if body == nil {
		return nil
	}
	return &BlindedBeaconBlockBodyCapella{
		RandaoReveal:           bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:               CopyETH1Data(body.Eth1Data),
		Graffiti:               bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings:      CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings:      CopyAttesterSlashings(body.AttesterSlashings),
		Attestations:           CopyAttestations(body.Attestations),
		Deposits:               CopyDeposits(body.Deposits),
		VoluntaryExits:         CopySignedVoluntaryExits(body.VoluntaryExits),
		SyncAggregate:          CopySyncAggregate(body.SyncAggregate),
		ExecutionPayloadHeader: body.ExecutionPayloadHeader.Copy(),
		BlsToExecutionChanges:  CopyBLSToExecutionChanges(body.BlsToExecutionChanges),
	}
}

// CopySignedBlindedBeaconBlockDeneb copies the provided SignedBlindedBeaconBlockDeneb.
func CopySignedBlindedBeaconBlockDeneb(sigBlock *SignedBlindedBeaconBlockDeneb) *SignedBlindedBeaconBlockDeneb {
	if sigBlock == nil {
		return nil
	}
	return &SignedBlindedBeaconBlockDeneb{
		Message:   CopyBlindedBeaconBlockDeneb(sigBlock.Message),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBlindedBeaconBlockDeneb copies the provided BlindedBeaconBlockDeneb.
func CopyBlindedBeaconBlockDeneb(block *BlindedBeaconBlockDeneb) *BlindedBeaconBlockDeneb {
	if block == nil {
		return nil
	}
	return &BlindedBeaconBlockDeneb{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBlindedBeaconBlockBodyDeneb(block.Body),
	}
}

// CopyBlindedBeaconBlockBodyDeneb copies the provided BlindedBeaconBlockBodyDeneb.
func CopyBlindedBeaconBlockBodyDeneb(body *BlindedBeaconBlockBodyDeneb) *BlindedBeaconBlockBodyDeneb {
	if body == nil {
		return nil
	}
	return &BlindedBeaconBlockBodyDeneb{
		RandaoReveal:           bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:               CopyETH1Data(body.Eth1Data),
		Graffiti:               bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings:      CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings:      CopyAttesterSlashings(body.AttesterSlashings),
		Attestations:           CopyAttestations(body.Attestations),
		Deposits:               CopyDeposits(body.Deposits),
		VoluntaryExits:         CopySignedVoluntaryExits(body.VoluntaryExits),
		SyncAggregate:          CopySyncAggregate(body.SyncAggregate),
		ExecutionPayloadHeader: body.ExecutionPayloadHeader.Copy(),
		BlsToExecutionChanges:  CopyBLSToExecutionChanges(body.BlsToExecutionChanges),
		BlobKzgCommitments:     CopyBlobKZGs(body.BlobKzgCommitments),
	}
}

// CopySignedBlindedBeaconBlockElectra copies the provided SignedBlindedBeaconBlockElectra.
func CopySignedBlindedBeaconBlockElectra(sigBlock *SignedBlindedBeaconBlockElectra) *SignedBlindedBeaconBlockElectra {
	if sigBlock == nil {
		return nil
	}
	return &SignedBlindedBeaconBlockElectra{
		Message:   CopyBlindedBeaconBlockElectra(sigBlock.Message),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBlindedBeaconBlockElectra copies the provided BlindedBeaconBlockElectra.
func CopyBlindedBeaconBlockElectra(block *BlindedBeaconBlockElectra) *BlindedBeaconBlockElectra {
	if block == nil {
		return nil
	}
	return &BlindedBeaconBlockElectra{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBlindedBeaconBlockBodyElectra(block.Body),
	}
}

// CopyBlindedBeaconBlockBodyElectra copies the provided BlindedBeaconBlockBodyElectra.
func CopyBlindedBeaconBlockBodyElectra(body *BlindedBeaconBlockBodyElectra) *BlindedBeaconBlockBodyElectra {
	if body == nil {
		return nil
	}
	return &BlindedBeaconBlockBodyElectra{
		RandaoReveal:           bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:               CopyETH1Data(body.Eth1Data),
		Graffiti:               bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings:      CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings:      CopyAttesterSlashingsElectra(body.AttesterSlashings),
		Attestations:           CopyAttestationsElectra(body.Attestations),
		Deposits:               CopyDeposits(body.Deposits),
		VoluntaryExits:         CopySignedVoluntaryExits(body.VoluntaryExits),
		SyncAggregate:          CopySyncAggregate(body.SyncAggregate),
		ExecutionPayloadHeader: body.ExecutionPayloadHeader.Copy(),
		BlsToExecutionChanges:  CopyBLSToExecutionChanges(body.BlsToExecutionChanges),
		BlobKzgCommitments:     CopyBlobKZGs(body.BlobKzgCommitments),
	}
}

// CopySignedBlindedBeaconBlockBellatrix copies the provided SignedBlindedBeaconBlockBellatrix.
func CopySignedBlindedBeaconBlockBellatrix(sigBlock *SignedBlindedBeaconBlockBellatrix) *SignedBlindedBeaconBlockBellatrix {
	if sigBlock == nil {
		return nil
	}
	return &SignedBlindedBeaconBlockBellatrix{
		Block:     CopyBlindedBeaconBlockBellatrix(sigBlock.Block),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBlindedBeaconBlockBellatrix copies the provided BlindedBeaconBlockBellatrix.
func CopyBlindedBeaconBlockBellatrix(block *BlindedBeaconBlockBellatrix) *BlindedBeaconBlockBellatrix {
	if block == nil {
		return nil
	}
	return &BlindedBeaconBlockBellatrix{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBlindedBeaconBlockBodyBellatrix(block.Body),
	}
}

// CopyBlindedBeaconBlockBodyBellatrix copies the provided BlindedBeaconBlockBodyBellatrix.
func CopyBlindedBeaconBlockBodyBellatrix(body *BlindedBeaconBlockBodyBellatrix) *BlindedBeaconBlockBodyBellatrix {
	if body == nil {
		return nil
	}
	return &BlindedBeaconBlockBodyBellatrix{
		RandaoReveal:           bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:               CopyETH1Data(body.Eth1Data),
		Graffiti:               bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings:      CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings:      CopyAttesterSlashings(body.AttesterSlashings),
		Attestations:           CopyAttestations(body.Attestations),
		Deposits:               CopyDeposits(body.Deposits),
		VoluntaryExits:         CopySignedVoluntaryExits(body.VoluntaryExits),
		SyncAggregate:          CopySyncAggregate(body.SyncAggregate),
		ExecutionPayloadHeader: body.ExecutionPayloadHeader.Copy(),
	}
}

func CopyBLSToExecutionChanges(changes []*SignedBLSToExecutionChange) []*SignedBLSToExecutionChange {
	if changes == nil {
		return nil
	}

	res := make([]*SignedBLSToExecutionChange, len(changes))
	for i := 0; i < len(changes); i++ {
		res[i] = &SignedBLSToExecutionChange{
			Message: &BLSToExecutionChange{
				ValidatorIndex:     changes[i].Message.ValidatorIndex,
				FromBlsPubkey:      bytesutil.SafeCopyBytes(changes[i].Message.FromBlsPubkey),
				ToExecutionAddress: bytesutil.SafeCopyBytes(changes[i].Message.ToExecutionAddress),
			},
			Signature: bytesutil.SafeCopyBytes(changes[i].Signature),
		}
	}

	return res
}

// CopyBlobKZGs copies the provided blob kzgs object.
func CopyBlobKZGs(b [][]byte) [][]byte {
	return bytesutil.SafeCopy2dBytes(b)
}

// CopySignedBeaconBlockDeneb copies the provided SignedBeaconBlockDeneb.
func CopySignedBeaconBlockDeneb(sigBlock *SignedBeaconBlockDeneb) *SignedBeaconBlockDeneb {
	if sigBlock == nil {
		return nil
	}
	return &SignedBeaconBlockDeneb{
		Block:     CopyBeaconBlockDeneb(sigBlock.Block),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBeaconBlockDeneb copies the provided BeaconBlockDeneb.
func CopyBeaconBlockDeneb(block *BeaconBlockDeneb) *BeaconBlockDeneb {
	if block == nil {
		return nil
	}
	return &BeaconBlockDeneb{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBeaconBlockBodyDeneb(block.Body),
	}
}

// CopyBeaconBlockBodyDeneb copies the provided BeaconBlockBodyDeneb.
func CopyBeaconBlockBodyDeneb(body *BeaconBlockBodyDeneb) *BeaconBlockBodyDeneb {
	if body == nil {
		return nil
	}
	return &BeaconBlockBodyDeneb{
		RandaoReveal:          bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:              CopyETH1Data(body.Eth1Data),
		Graffiti:              bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings:     CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings:     CopyAttesterSlashings(body.AttesterSlashings),
		Attestations:          CopyAttestations(body.Attestations),
		Deposits:              CopyDeposits(body.Deposits),
		VoluntaryExits:        CopySignedVoluntaryExits(body.VoluntaryExits),
		SyncAggregate:         CopySyncAggregate(body.SyncAggregate),
		ExecutionPayload:      body.ExecutionPayload.Copy(),
		BlsToExecutionChanges: CopyBLSToExecutionChanges(body.BlsToExecutionChanges),
		BlobKzgCommitments:    CopyBlobKZGs(body.BlobKzgCommitments),
	}
}

// CopySignedBeaconBlockElectra copies the provided SignedBeaconBlockElectra.
func CopySignedBeaconBlockElectra(sigBlock *SignedBeaconBlockElectra) *SignedBeaconBlockElectra {
	if sigBlock == nil {
		return nil
	}
	return &SignedBeaconBlockElectra{
		Block:     CopyBeaconBlockElectra(sigBlock.Block),
		Signature: bytesutil.SafeCopyBytes(sigBlock.Signature),
	}
}

// CopyBeaconBlockElectra copies the provided BeaconBlockElectra.
func CopyBeaconBlockElectra(block *BeaconBlockElectra) *BeaconBlockElectra {
	if block == nil {
		return nil
	}
	return &BeaconBlockElectra{
		Slot:          block.Slot,
		ProposerIndex: block.ProposerIndex,
		ParentRoot:    bytesutil.SafeCopyBytes(block.ParentRoot),
		StateRoot:     bytesutil.SafeCopyBytes(block.StateRoot),
		Body:          CopyBeaconBlockBodyElectra(block.Body),
	}
}

// CopyBeaconBlockBodyElectra copies the provided BeaconBlockBodyElectra.
func CopyBeaconBlockBodyElectra(body *BeaconBlockBodyElectra) *BeaconBlockBodyElectra {
	if body == nil {
		return nil
	}
	return &BeaconBlockBodyElectra{
		RandaoReveal:          bytesutil.SafeCopyBytes(body.RandaoReveal),
		Eth1Data:              CopyETH1Data(body.Eth1Data),
		Graffiti:              bytesutil.SafeCopyBytes(body.Graffiti),
		ProposerSlashings:     CopyProposerSlashings(body.ProposerSlashings),
		AttesterSlashings:     CopyAttesterSlashingsElectra(body.AttesterSlashings),
		Attestations:          CopyAttestationsElectra(body.Attestations),
		Deposits:              CopyDeposits(body.Deposits),
		VoluntaryExits:        CopySignedVoluntaryExits(body.VoluntaryExits),
		SyncAggregate:         CopySyncAggregate(body.SyncAggregate),
		ExecutionPayload:      body.ExecutionPayload.Copy(),
		BlsToExecutionChanges: CopyBLSToExecutionChanges(body.BlsToExecutionChanges),
		BlobKzgCommitments:    CopyBlobKZGs(body.BlobKzgCommitments),
	}
}

// CopyHistoricalSummaries copies the historical summaries.
func CopyHistoricalSummaries(summaries []*HistoricalSummary) []*HistoricalSummary {
	if summaries == nil {
		return nil
	}
	newSummaries := make([]*HistoricalSummary, len(summaries))
	for i, s := range summaries {
		newSummaries[i] = &HistoricalSummary{
			BlockSummaryRoot: bytesutil.SafeCopyBytes(s.BlockSummaryRoot),
			StateSummaryRoot: bytesutil.SafeCopyBytes(s.StateSummaryRoot),
		}
	}
	return newSummaries
}

// CopyPartialWithdrawals copies the provided partial withdrawals.
func CopyPendingPartialWithdrawals(pws []*PendingPartialWithdrawal) []*PendingPartialWithdrawal {
	if pws == nil {
		return nil
	}
	newPws := make([]*PendingPartialWithdrawal, len(pws))
	for i, pw := range pws {
		newPws[i] = &PendingPartialWithdrawal{
			Index:             pw.Index,
			Amount:            pw.Amount,
			WithdrawableEpoch: pw.WithdrawableEpoch,
		}
	}
	return newPws
}

func CopyPendingConsolidations(pcs []*PendingConsolidation) []*PendingConsolidation {
	if pcs == nil {
		return nil
	}
	newPcs := make([]*PendingConsolidation, len(pcs))
	for i, pc := range pcs {
		newPcs[i] = &PendingConsolidation{
			SourceIndex: pc.SourceIndex,
			TargetIndex: pc.TargetIndex,
		}
	}
	return newPcs
}

func CopyPendingBalanceDeposits(pbd []*PendingBalanceDeposit) []*PendingBalanceDeposit {
	if pbd == nil {
		return nil
	}
	newPbd := make([]*PendingBalanceDeposit, len(pbd))
	for i, pb := range pbd {
		newPbd[i] = &PendingBalanceDeposit{
			Index:  pb.Index,
			Amount: pb.Amount,
		}
	}
	return newPbd
}
