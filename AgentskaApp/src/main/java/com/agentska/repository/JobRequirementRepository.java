package com.agentska.repository;

import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.agentska.model.JobRequirement;

@Repository
public interface JobRequirementRepository extends JpaRepository<JobRequirement, Long> {
	Optional<JobRequirement> findByJobId(Integer job);
}
